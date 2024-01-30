import { writable } from "svelte/store";
import { PatchRequest, DeleteRequest } from "./wailsjs/go/main/App";
import Swal from "sweetalert2";

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

        const response = await PatchRequest(url, JSON.stringify(formData))

        if(!response.success) {
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
    delete: async function() {
        const url = `https://api.myanimelist.net/v2/anime/${this.animeId}/my_list_status`

        this.close()

        const result = await Swal.fire({
            title: "Are you sure?",
            text: "You won't be able to revert this!",
            icon: "warning",
            showCancelButton: true,
            confirmButtonText: "Yes, delete it!",
            cancelButtonText: "No, cancel!",  
        }).then((result) => {
            return result.isConfirmed
        })

        if(!result) {
            return
        }

        const response = await DeleteRequest(url)

        if(!response.success) {
            console.log("error")
            return
        }

        ListItems.update((list) => {
            list = list.filter((item) => {
                return item.node.id !== this.animeId
            })

            return list
        })
    },
    close: function () {
        this.element.classList.add("inactive")
    },
    open: function () {
        this.element.classList.remove("inactive")
    }
});