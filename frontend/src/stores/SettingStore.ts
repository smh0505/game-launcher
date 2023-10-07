import { defineStore } from "pinia"

export const useSettingStore = defineStore("setting", {
    state: () => ({
        chroma: 0,
        hue: 0,
        opacity: 0.5,
        background: ""
    }),
    actions: {
        setColor() {
            document.documentElement.style.setProperty("--chroma", this.chroma.toString())
            document.documentElement.style.setProperty("--hue", this.hue.toString())
            document.documentElement.style.setProperty("--opacity", this.opacity.toString())
        },
        saveSetting() {
            localStorage.setItem("setting", JSON.stringify({
                chroma: this.chroma,
                hue: this.hue,
                opacity: this.opacity,
                background: this.background
            }))
        },
        loadSetting() {
            const setting = localStorage.getItem("setting")
            if (setting) {
                const group = JSON.parse(setting)
                this.chroma = group.chroma ?? 0
                this.hue = group.hue ?? 0
                this.opacity = group.opacity ?? 0.5
                this.background = group.background ? group.background : ""
                this.setColor()
            }
        }
    }
})
