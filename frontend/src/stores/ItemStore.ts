import { defineStore } from "pinia"
import { useSettingStore } from "./SettingStore.js"
import { SaveMetadata, LoadMetadata } from "../../wailsjs/go/main/App.js"

export const useItemStore = defineStore('items', {
    state: () => ({
        items: [] as Array<Game>
    }),
    actions: {
        addItem(path: string, name: string, link: string) {
            this.items.push({
                id: this.items.length,
                image: "",
                code: "",
                path,
                name,
                link,
                isPlayed: IsPlayed.notPlayed
            })
        },
        checkUnique(name: string) {
            return this.items.filter(x => x.name === name).length < 2
        },
        sortItems(type: 'name' | 'code' | 'isPlayed', reverse: boolean = false) {
            switch (type) {
                case "name":
                    this.items.sort((a, b) => reverse ? b.name.localeCompare(a.name) : a.name.localeCompare(b.name))
                    break
                
                case "code":
                    const codes = this.items
                        .filter(item => item.code.trim() !== "")
                        .toSorted((a, b) => reverse ? b.code.localeCompare(a.code) : a.code.localeCompare(b.code))
                    const nonCodes = this.items
                        .filter(item => item.code.trim() === "")
                        .toSorted((a, b) => reverse ? b.name.localeCompare(a.name) : a.name.localeCompare(b.name))
                    
                    this.items = reverse ? [...nonCodes, ...codes] : [...codes, ...nonCodes] 
                    break
                
                case "isPlayed":
                    const cleared = partialSort(this.items, IsPlayed.cleared, reverse)
                    const played = partialSort(this.items, IsPlayed.played, reverse)
                    const notPlayed = partialSort(this.items, IsPlayed.notPlayed, reverse)
                    
                    this.items = reverse ? [...notPlayed, ...played, ...cleared] : [...cleared, ...played, ...notPlayed]
            }
            this.items.forEach((x, id) => x.id = id)
            useSettingStore().isSaving = true
            this.saveItems()
        },
        setIfPlayed(index: number) {
            this.items[index].isPlayed = (this.items[index].isPlayed + 1) % 3
            this.saveItems()
        },
        saveItems() {
            SaveMetadata(this.items).then(() => useSettingStore().isSaving = false)
        },
        loadItems() {
            LoadMetadata().then(list => {
                if (list) {
                    for (let item of list) {
                        this.items.push({
                            id: this.items.length,
                            image: item.image ?? "",
                            code: item.code ?? "",
                            path: item.path ?? "",
                            name: item.name ?? "",
                            link: item.link ?? "",
                            isPlayed: item.isPlayed ?? IsPlayed.notPlayed
                        })
                    }
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
    link: string,
    isPlayed: IsPlayed
}

enum IsPlayed {
    notPlayed = 0,
    played,
    cleared
}

const partialSort = (items: Game[], isPlayed: IsPlayed, reverse: boolean) => {
    return items
        .filter(item => item.isPlayed === isPlayed)
        .toSorted((a, b) => reverse ? b.name.localeCompare(a.name) : a.name.localeCompare(b.name))
}
