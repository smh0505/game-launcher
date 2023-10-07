import { defineStore } from "pinia"
import { SaveMetadata, LoadMetadata } from "../../wailsjs/go/main/App.js"

export const useItemStore = defineStore('items', {
    state: () => ({
        items: [] as Item[]
    }),
    actions: {
        addItem(path: string, name: string, link: string) {
            this.items.push({ image: "", path, name, link })
        },
        saveItems() {
            SaveMetadata(this.items)
        },
        loadItems() {
            LoadMetadata().then(x => this.items = structuredClone(x)).then(_ => console.log(this.items))
        }
    }
})

interface Item {
    image: string,
    path: string,
    name: string,
    link: string
}
