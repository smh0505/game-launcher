import { defineStore } from "pinia"
import { useSettingStore } from "./SettingStore.js"
import { SaveMetadata, LoadMetadata } from "../../wailsjs/go/main/App.js"

export const useItemStore = defineStore('items', {
    state: () => ({
        items: [] as Array<Game>
    }),
    actions: {
        addItem(path: string, name: string, link: string) {
            this.items.push({ id: this.items.length, image: "", code: "", path, name, link })
        },
        checkUnique(name: string) {
            return this.items.filter(x => x.name === name).length < 2
        },
        sortItems(type: 'name' | 'code', reverse: boolean = false) {
            this.items.sort((a, b) => {
                switch (type) {
                    case "name":
                        return reverse ? b.name.localeCompare(a.name) : a.name.localeCompare(b.name)
                    case "code":
                        return reverse ? b.code.localeCompare(a.code) : a.code.localeCompare(b.code)
                }
            })
            this.items.forEach((x, id) => x.id = id)
        },
        saveItems() {
            SaveMetadata(this.items).then(() => useSettingStore().isSaving = false)
        },
        loadItems() {
            LoadMetadata().then(x => {
                if (x) {
                    this.items = structuredClone(x)
                    this.items.forEach(item => {
                        const empty = {
                            id: 0,
                            image: "",
                            code: "",
                            path: "",
                            name: "",
                            link: ""
                        }
                        item = { ...empty, ...item }
                    })
                }
            })
        }
    }
})

export type Game = {
    id: number,
    image: string,
    code: string,
    path: string,
    name: string,
    link: string
}
