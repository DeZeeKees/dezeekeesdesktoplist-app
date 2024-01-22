import { writable } from "svelte/store";
import { PatchRequest } from "./wailsjs/go/main/App";

export let title = writable("");
export let animeCurrentTab = writable("watching");

export let ListItems = writable(undefined);

export let popover = writable({
    element: undefined,
    title: "",
    episodesWatched: 0,
    episodesTotal: 0,
    status: "",
    score: 0,
    rewatching: false,
    timesRewatched: 0,
    animeId: 0,
    save: async function () {
        const url = `https://api.myanimelist.net/v2/anime/${this.animeId}/my_list_status`

        const formData = {
            num_watched_episodes: this.episodesWatched,
            score: this.score,
            status: this.status,
            is_rewatching: this.rewatching,
            num_times_rewatched: this.timesRewatched
        }

        const response = await PatchRequest(url, JSON.stringify(formData)).then((data) => {
            if(data === "error") {
                console.log("error")
                return
            }

            return JSON.parse(data)
        })

        if(response === undefined || response === "error") {
            console.log("error")
            return
        }

        ListItems.update((list) => {
            const index = list.findIndex((item) => {
                return item.node.id === this.animeId
            })

            let currentTab = ""
            animeCurrentTab.subscribe($ => currentTab = $)()

            if(this.status !== currentTab) {
                list = list.filter((item) => {
                    return item.node.id !== this.animeId
                })
                return list
            }

            list[index].list_status.num_episodes_watched = this.episodesWatched
            list[index].list_status.score = this.score
            list[index].list_status.status = this.status
            list[index].list_status.is_rewatching = this.rewatching
            list[index].list_status.num_times_rewatched = this.timesRewatched

            return list
        })

        this.close()
    },
    close: function () {
        this.element.classList.add("inactive")
    },
    open: function () {
        this.element.classList.remove("inactive")
    }
});