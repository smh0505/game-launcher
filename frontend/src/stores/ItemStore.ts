import { defineStore } from "pinia"
import { SaveMetadata, LoadMetadata } from "../../wailsjs/go/main/App.js"

export const useItemStore = defineStore('items', {
    state: () => ({
        items: [] as Array<Game>
    }),
    actions: {
        addItem(path: string, name: string, link: string) {
            this.items.push({ image: "", path: path, name: name, link: link })
        },
        checkUnique(name: string) {
            return this.items.filter(x => x.name === name).length === 1
        },
        saveItems() {
            SaveMetadata(this.items)
        },
        loadItems() {
            LoadMetadata().then(x => { if (x) {this.items = structuredClone(x)}})
        }
    }
})

type Game = {
    image: string,
    path: string,
    name: string,
    link: string
}
