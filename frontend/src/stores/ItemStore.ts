import { defineStore } from "pinia"

export const useItemStore = defineStore('items', {
    state: () => ({
        items: [] as Item[]
    }),
    actions: {
        addItem(path: string, name: string, link: string) {
            this.items.push({ path, name, link })
        }
    }
})

interface Item {
    path: string,
    name: string,
    link: string
}
