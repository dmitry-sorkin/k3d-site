/**
 * ProfileManager - модульная система управления профилями для калибровщиков
 * Позволяет сохранять, загружать, экспортировать и импортировать профили настроек
 */
class ProfileManager {
    constructor(storagePrefix = 'k3d_profiles', maxProfiles = 10) {
        this.storagePrefix = storagePrefix;
        this.maxProfiles = maxProfiles;
        this.currentProfileId = null;
        this.profiles = this.loadProfiles();
        this.callbacks = {
            onProfileChange: null,
            onProfileSave: null,
            onProfileDelete: null
        };
    }

    /**
     * Загрузка всех профилей из localStorage
     */
    loadProfiles() {
        const stored = localStorage.getItem(`${this.storagePrefix}_list`);
        if (!stored) {
            return [];
        }
        try {
            return JSON.parse(stored);
        } catch (e) {
            console.error('Ошибка загрузки профилей:', e);
            return [];
        }
    }

    /**
     * Сохранение списка профилей в localStorage
     */
    saveProfiles() {
        localStorage.setItem(`${this.storagePrefix}_list`, JSON.stringify(this.profiles));
    }

    /**
     * Создание нового профиля
     */
    createProfile(name, settings = {}) {
        if (this.profiles.length >= this.maxProfiles) {
            throw new Error(`Достигнут лимит профилей (максимум ${this.maxProfiles})`);
        }

        const profile = {
            id: this.generateId(),
            name: name || `Профиль ${this.profiles.length + 1}`,
            createdAt: Date.now(),
            updatedAt: Date.now(),
            settings: settings
        };

        this.profiles.push(profile);
        this.saveProfiles();

        if (this.callbacks.onProfileSave) {
            this.callbacks.onProfileSave(profile);
        }

        return profile;
    }

    /**
     * Обновление существующего профиля
     */
    updateProfile(profileId, settings) {
        const profile = this.profiles.find(p => p.id === profileId);
        if (!profile) {
            throw new Error('Профиль не найден');
        }

        profile.settings = settings;
        profile.updatedAt = Date.now();
        this.saveProfiles();

        if (this.callbacks.onProfileSave) {
            this.callbacks.onProfileSave(profile);
        }

        return profile;
    }

    /**
     * Обновление имени профиля
     */
    renameProfile(profileId, newName) {
        const profile = this.profiles.find(p => p.id === profileId);
        if (!profile) {
            throw new Error('Профиль не найден');
        }

        profile.name = newName;
        profile.updatedAt = Date.now();
        this.saveProfiles();

        return profile;
    }

    /**
     * Удаление профиля
     */
    deleteProfile(profileId) {
        const index = this.profiles.findIndex(p => p.id === profileId);
        if (index === -1) {
            throw new Error('Профиль не найден');
        }

        const deleted = this.profiles.splice(index, 1)[0];
        this.saveProfiles();

        if (this.callbacks.onProfileDelete) {
            this.callbacks.onProfileDelete(deleted);
        }

        // Если удален текущий профиль, сбрасываем выбор
        if (this.currentProfileId === profileId) {
            this.currentProfileId = null;
        }

        return deleted;
    }

    /**
     * Получение профиля по ID
     */
    getProfile(profileId) {
        return this.profiles.find(p => p.id === profileId);
    }

    /**
     * Получение всех профилей
     */
    getAllProfiles() {
        return [...this.profiles];
    }

    /**
     * Установка текущего активного профиля
     */
    setCurrentProfile(profileId) {
        const profile = this.getProfile(profileId);
        if (!profile && profileId !== null) {
            throw new Error('Профиль не найден');
        }

        this.currentProfileId = profileId;
        localStorage.setItem(`${this.storagePrefix}_current`, profileId || '');

        if (this.callbacks.onProfileChange) {
            this.callbacks.onProfileChange(profile);
        }

        return profile;
    }

    /**
     * Получение текущего профиля
     */
    getCurrentProfile() {
        if (!this.currentProfileId) {
            // Попытка загрузить из localStorage
            const stored = localStorage.getItem(`${this.storagePrefix}_current`);
            if (stored) {
                this.currentProfileId = stored;
            }
        }
        return this.getProfile(this.currentProfileId);
    }

    /**
     * Экспорт профиля в JSON
     */
    exportProfile(profileId) {
        const profile = this.getProfile(profileId);
        if (!profile) {
            throw new Error('Профиль не найден');
        }

        const exportData = {
            version: '1.0',
            exported: Date.now(),
            profile: profile
        };

        return JSON.stringify(exportData, null, 2);
    }

    /**
     * Экспорт всех профилей
     */
    exportAllProfiles() {
        const exportData = {
            version: '1.0',
            exported: Date.now(),
            profiles: this.profiles
        };

        return JSON.stringify(exportData, null, 2);
    }

    /**
     * Импорт профиля из JSON
     */
    importProfile(jsonString, overwrite = false) {
        try {
            const data = JSON.parse(jsonString);

            // Проверка версии
            if (!data.version || !data.profile) {
                throw new Error('Неверный формат файла профиля');
            }

            const profile = data.profile;

            // Проверка на дубликат
            const existing = this.profiles.find(p => p.id === profile.id);
            if (existing && !overwrite) {
                // Создаем новый ID для импортированного профиля
                profile.id = this.generateId();
                profile.name = `${profile.name} (импорт)`;
            }

            // Обновление временных меток
            profile.updatedAt = Date.now();

            if (existing && overwrite) {
                // Перезаписываем существующий профиль
                const index = this.profiles.findIndex(p => p.id === profile.id);
                this.profiles[index] = profile;
            } else {
                // Добавляем как новый профиль
                if (this.profiles.length >= this.maxProfiles) {
                    throw new Error(`Достигнут лимит профилей (максимум ${this.maxProfiles})`);
                }
                this.profiles.push(profile);
            }

            this.saveProfiles();
            return profile;

        } catch (e) {
            throw new Error(`Ошибка импорта профиля: ${e.message}`);
        }
    }

    /**
     * Импорт нескольких профилей
     */
    importAllProfiles(jsonString, overwrite = false) {
        try {
            const data = JSON.parse(jsonString);

            if (!data.version || !data.profiles || !Array.isArray(data.profiles)) {
                throw new Error('Неверный формат файла профилей');
            }

            const imported = [];
            for (const profile of data.profiles) {
                try {
                    const importedProfile = this.importProfile(
                        JSON.stringify({ version: data.version, profile }),
                        overwrite
                    );
                    imported.push(importedProfile);
                } catch (e) {
                    console.error(`Ошибка импорта профиля ${profile.name}:`, e);
                }
            }

            return imported;

        } catch (e) {
            throw new Error(`Ошибка импорта профилей: ${e.message}`);
        }
    }

    /**
     * Очистка всех профилей
     */
    clearAllProfiles() {
        this.profiles = [];
        this.currentProfileId = null;
        this.saveProfiles();
        localStorage.removeItem(`${this.storagePrefix}_current`);
    }

    /**
     * Генерация уникального ID
     */
    generateId() {
        return `${Date.now()}-${Math.random().toString(36).substr(2, 9)}`;
    }

    /**
     * Установка callback функций
     */
    setCallbacks(callbacks) {
        Object.assign(this.callbacks, callbacks);
    }

    /**
     * Валидация настроек профиля
     * Может быть переопределена для конкретного калибровщика
     */
    validateSettings(settings) {
        // Базовая валидация - проверка что settings это объект
        if (!settings || typeof settings !== 'object') {
            throw new Error('Настройки должны быть объектом');
        }
        return true;
    }

    /**
     * Получение дефолтных настроек
     * Должна быть переопределена для конкретного калибровщика
     */
    getDefaultSettings() {
        return {};
    }
}