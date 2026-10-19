#!/usr/bin/env python3
"""Делает прямые ссылки на AliExpress аффилиатными через API epn.bz.

Ровно одна страница за запуск. Ссылки на другие маркетплейсы не трогаются вообще.
Уже аффилиатные ссылки (ali.click, alli.pub, ...) пропускаются — повторный прогон безопасен.

Запуск:  python tools/affiliate.py docs/part-navi/filament_dryers.md
Полное описание и разбор ошибок — в AGENTS.md, раздел «Аффилиатные ссылки epn.bz».
"""
from __future__ import annotations

import json
import os
import pathlib
import re
import subprocess
import sys
import time
import urllib.error
import urllib.parse
import urllib.request
from typing import NoReturn

API = "https://app.epn.bz"
OAUTH = "https://oauth2.epn.bz"

# AliExpress в EPN для сайта k3d.tech. Площадка привязана к офферу: с чужой площадкой
# API отвечает 422 «Offer does not match placement!». Менять только вместе с оффером.
ALIEXPRESS_OFFER_ID = 1
ALIEXPRESS_PLACEMENT_ID = 1455321

# /creative/create: лимит 1 запрос в секунду.
CREATIVE_PAUSE = 1.2

LINK_RE = re.compile(
    r"https?://(?:[a-z0-9-]+\.)*aliexpress\.(?:ru|com|us)/[^\s)\"'<>\]\[{}]+",
    re.IGNORECASE,
)

ROOT = pathlib.Path(__file__).resolve().parent.parent
TOKENS = ROOT / ".cache" / "epn_tokens.json"


def die(msg: str) -> NoReturn:
    sys.exit(f"ОШИБКА: {msg}")


def cred(name: str) -> str | None:
    """setx пишет в HKCU\\Environment, а уже запущенный процесс его не видит."""
    value = os.environ.get(name)
    if value:
        return value
    if sys.platform == "win32":
        out = subprocess.run(
            ["powershell", "-NoProfile", "-Command",
             f"[Environment]::GetEnvironmentVariable('{name}','User')"],
            capture_output=True, text=True)
        return out.stdout.strip() or None
    return None


class Epn:
    def __init__(self) -> None:
        self.cid = cred("EPN_CLIENT_ID")
        self.secret = cred("EPN_CLIENT_SECRET")
        if not (self.cid and self.secret):
            die("не найдены EPN_CLIENT_ID / EPN_CLIENT_SECRET.\n"
                "Задай их так (и открой новое окно терминала):\n"
                '  setx EPN_CLIENT_ID "..."\n'
                '  setx EPN_CLIENT_SECRET "..."')
        self.tok = self._load_tokens()

    @staticmethod
    def _raw(url: str, method: str = "GET", body=None, token=None,
             api_version: bool = False) -> tuple[int, dict]:
        # Все адреса — только жёстко заданные константы API, но схему проверяем явно.
        if not url.startswith("https://"):
            die(f"недопустимый адрес запроса: {url}")
        headers = {"ACCEPT-LANGUAGE": "ru"}
        if api_version:
            headers["X-API-VERSION"] = "2"
        if token:
            headers["X-ACCESS-TOKEN"] = token
        data = None
        if body is not None:
            data = json.dumps(body).encode()
            headers["Content-Type"] = "application/json"
        req = urllib.request.Request(url, data=data, headers=headers, method=method)
        try:
            with urllib.request.urlopen(req, timeout=40) as r:  # nosec B310 — схема проверена выше
                return r.status, json.loads(r.read().decode())
        except urllib.error.HTTPError as e:
            try:
                return e.code, json.loads(e.read().decode() or "{}")
            except ValueError:
                return e.code, {}
        except urllib.error.URLError as e:
            return 0, {"errors": [{"error_description": f"сеть недоступна: {e.reason}"}]}

    def _save(self, attrs: dict) -> dict:
        self.tok = {"access_token": attrs["access_token"],
                    "refresh_token": attrs.get("refresh_token", "")}
        TOKENS.parent.mkdir(exist_ok=True)
        TOKENS.write_text(json.dumps(self.tok, indent=2), encoding="utf-8")
        return self.tok

    def _load_tokens(self) -> dict:
        # Битый или устаревший кэш — не повод падать, просто берём новые токены.
        try:
            cached = json.loads(TOKENS.read_text(encoding="utf-8"))
        except (OSError, ValueError):
            cached = {}
        if cached.get("access_token"):
            return cached
        return self._new_tokens()

    def _new_tokens(self) -> dict:
        # Внимание: /ssid — 1 запрос в сутки с одного IP. Поэтому токены кэшируются.
        st, r = self._raw(f"{OAUTH}/ssid?v=2&client_id={self.cid}", api_version=True)
        if st != 200 or "data" not in r:
            die(f"не удалось получить ssid-токен: {st} {json.dumps(r, ensure_ascii=False)}\n"
                "Лимит метода — 1 запрос в сутки с одного IP; второй требует капчу.")
        st, r = self._raw(f"{OAUTH}/token?v=2", "POST", {
            "ssid_token": r["data"]["attributes"]["ssid_token"],
            "client_id": self.cid, "client_secret": self.secret,
            "grant_type": "client_credential", "check_ip": False,
        }, api_version=True)
        if st != 200:
            die(f"не удалось получить access_token: {st} {json.dumps(r, ensure_ascii=False)}")
        return self._save(r["data"]["attributes"])

    def _refresh(self) -> None:
        st, r = self._raw(f"{OAUTH}/token/refresh", "POST", {
            "grant_type": "refresh_token", "refresh_token": self.tok.get("refresh_token"),
            "client_id": self.cid,
        }, api_version=True)
        if st != 200:
            die(f"refresh_token не сработал: {st} {json.dumps(r, ensure_ascii=False)}\n"
                f"Удали {TOKENS.relative_to(ROOT)} и повтори — потребуется новый /ssid.")
        self._save(r["data"]["attributes"])

    def call(self, path: str, method: str = "GET", body=None, params=None) -> tuple[int, dict]:
        url = API + path + ("?" + urllib.parse.urlencode(params) if params else "")
        st, r = self._raw(url, method, body, token=self.tok["access_token"])
        if st == 401:
            self._refresh()
            st, r = self._raw(url, method, body, token=self.tok["access_token"])
        return st, r

    def offer_of(self, link: str) -> tuple[int, object, str]:
        st, r = self.call("/offers/by-link", params={"link": link})
        data = r.get("data") or {}
        attrs = data.get("attributes") or {}
        return st, data.get("id"), attrs.get("title") or ""

    def creative(self, link: str, description: str) -> tuple[str | None, str | None]:
        """Возвращает (готовая аффилиатная ссылка, текст ошибки)."""
        for attempt in range(4):
            st, r = self.call("/creative/create", "POST", {
                "link": link, "offerId": ALIEXPRESS_OFFER_ID,
                "placementId": ALIEXPRESS_PLACEMENT_ID,
                "description": description, "type": "deeplink",
            })
            if st == 429:
                time.sleep(2 * (attempt + 1))
                continue
            if st == 200:
                return r["data"]["attributes"].get("code"), None
            errs = "; ".join(e.get("error_description", "?") for e in r.get("errors") or [])
            return None, f"{st} {errs or json.dumps(r, ensure_ascii=False)}"
        return None, "429 — превышен лимит запросов"


def describe(text: str, rel: str, line: int) -> str:
    """Подпись креатива в ЛК epn.bz. Обязательное поле, максимум 100 символов."""
    m = re.search(r"^title:\s*(.+)$", text, re.MULTILINE)
    base = m.group(1).strip().strip("\"'") if m else pathlib.Path(rel).stem
    return f"{base} [{rel}:{line}]"[:100]


def main(argv: list[str]) -> int:
    args = [a for a in argv if a != "--"]
    if len(args) != 1:
        die("нужен ровно один аргумент — путь к одной странице.\n"
            "  python tools/affiliate.py docs/part-navi/filament_dryers.md\n"
            "Скрипт намеренно не умеет обрабатывать несколько файлов или каталог.")
    target = pathlib.Path(args[0])
    if not target.is_file():
        die(f"«{target}» — не файл. Скрипт обрабатывает одну страницу за запуск.")
    try:
        rel = target.resolve().relative_to(ROOT / "docs").as_posix()
    except ValueError:
        die(f"«{target}» вне каталога docs/.")

    with target.open(encoding="utf-8", newline="") as fh:  # newline="" — сохранить CRLF как есть
        text = fh.read()

    found: list[tuple[str, int]] = []
    for m in LINK_RE.finditer(text):
        url = m.group(0).rstrip(".,;:")
        if "/aff/redirect" in url:
            continue
        found.append((url, text.count("\n", 0, m.start()) + 1))

    if not found:
        print(f"{rel}: прямых ссылок на AliExpress нет — файл не изменён.")
        return 0

    unique: dict[str, int] = {}
    for url, line in found:
        unique.setdefault(url, line)

    print(f"Файл: docs/{rel}")
    print(f"Ссылок на AliExpress: {len(found)} (уникальных: {len(unique)})\n")

    epn = Epn()
    replaced: dict[str, str] = {}
    errors = 0

    for url, line in sorted(unique.items(), key=lambda kv: -len(kv[0])):
        st, offer_id, offer_name = epn.offer_of(url)
        if st != 200:
            errors += 1
            print(f"  ✖ строка {line}  {url}\n      оффер не определён ({st}) — пропущено")
            continue
        if str(offer_id) != str(ALIEXPRESS_OFFER_ID):
            errors += 1
            print(f"  ✖ строка {line}  {url}\n"
                  f"      EPN отдал оффер {offer_id} «{offer_name}», а не AliExpress — пропущено")
            continue

        code, err = epn.creative(url, describe(text, rel, line))
        if err or not code:
            errors += 1
            print(f"  ✖ строка {line}  {url}\n      {err}")
            continue

        replaced[url] = code
        print(f"  ✔ строка {line}  {url}\n      → {code}")
        time.sleep(CREATIVE_PAUSE)

    for url in sorted(replaced, key=len, reverse=True):
        text = text.replace(url, replaced[url])

    if replaced:
        with target.open("w", encoding="utf-8", newline="") as fh:
            fh.write(text)
        print(f"\nЗаменено: {len(replaced)}. Файл docs/{rel} перезаписан.")
    else:
        print("\nНечего заменять — файл не изменён.")
    if errors:
        print(f"Не обработано: {errors} (см. выше). Остальные ссылки в файле не тронуты.")
    return 1 if errors else 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
