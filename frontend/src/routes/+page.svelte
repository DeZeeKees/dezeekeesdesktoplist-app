<Tabs tabs={tabList} on:tabChange={handleTabChange} />

<main bind:this={itemContainer} on:scroll={handleScroll}>
    {#if $ListItems === undefined}
        <h1>Loading...</h1>
    {:else if $ListItems === null}
        <h1>Failed to load data</h1>
    {:else}
        {#each $ListItems as item}
            <YourListItem data={item} bind:this={item.element} />
        {/each}
    {/if}
</main>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="popover inactive" bind:this={$popover.element} on:mouseup={closePopover()}>
    <div class="popover-content">
        <!-- svelte-ignore a11y-missing-content -->
        <h3>{$popover.title}</h3>

        <div class="options">
            <p>Status</p>
            <select bind:value={$popover.status}>
                <option value="watching">Watching</option>
                <option value="completed">Completed</option>
                <option value="on_hold">On Hold</option>
                <option value="dropped">Dropped</option>
                <option value="plan_to_watch">Plan to Watch</option>
            </select>
            <p>Episodes Watched</p>
            <div class="inline">
                <input type="number" bind:value={$popover.episodesWatched}> 
                <p>/ {$popover.episodesTotal}</p>
            </div>
            <p>Score</p>
            <select bind:value={$popover.score}>
                {#each [0,1,2,3,4,5,6,7,8,9,10] as score}
                    <option value={score}>{score}</option>
                {/each}
            </select>
            <p>Rewatching</p>
            <select bind:value={$popover.rewatching}>
                <option value="false">No</option>
                <option value="true">Yes</option>
            </select>
            <p>Times Rewatched</p>
            <input type="number" bind:value={$popover.timesRewatched}>
        </div>
        
        <div class="buttons">
            <button on:mouseup={() => { $popover.element.classList.add("inactive") }}>Close</button>
            <button on:click={$popover.delete()}>Delete</button>
            <button class="" on:click={$popover.save()}>Save</button>
        </div>
    </div>
</div>

<script>
    // @ts-nocheck

    import { title, popover, ListItems, animeCurrentTab } from "$lib/store";
    import { GetCurrentUser, GetRequest, GetSettings } from "$lib/wailsjs/go/main/App";
    import Tabs from "$lib/components/tabs.svelte";
    import YourListItem from "$lib/components/YourListItem.svelte";
    import * as styleManager from "$lib/styleManager";
    import { Toast } from "$lib";
    import { EventsOn } from "$lib/wailsjs/runtime/runtime";
    import { onMount } from "svelte";
    import { afterNavigate } from "$app/navigation";

    const tabList = [
        {name: "Watching"},
        {name: "Completed"},
        {name: "On Hold"},
        {name: "Dropped"},
        {name: "All"},
        {name: "Plan to Watch"}
    ]
    
    let ListResponse = undefined
    let itemContainer

    let canRefresh = true

    let settings = {
        usePrerelease: false,
        yourListCardSizeMultiplier: 1,
        nsfwContent: false,
    }

    const fields = "list_status{num_times_rewatched},num_episodes,media_type,start_date,end_date"
    const fetchInterval = 30

    let baseUrl = "https://api.myanimelist.net/v2/users/@me/animelist?limit=" + fetchInterval + "&fields=" + fields

    onMount(async () => {
        EventsOn("startup:complete", async () => {
            settings = await GetSettings()
            styleManager.set("your-list-card-size-multiplier", settings.yourListCardSizeMultiplier)

            baseUrl += "&nsfw=" + settings.nsfwContent

            const MalUser = await GetCurrentUser()

            title.set("Your List - " + MalUser.name)

            ListResponse = await GetStuff(baseUrl + "&status=watching")
            $ListItems = ListResponse.data

            console.log(settings.nsfwContent)
        })
    })

    afterNavigate(async () => {
        const MalUser = await GetCurrentUser()
        title.set("Your List - " + MalUser.name)

        settings = await GetSettings()
        styleManager.set("your-list-card-size-multiplier", settings.yourListCardSizeMultiplier)

        baseUrl += "&nsfw=" + settings.nsfwContent
    })

    /**
     * @param event {CustomEvent}
     */
    async function handleTabChange(event) {
        itemContainer.scrollTo(0, 0)
        $ListItems = undefined

        if(event.detail === "all") {
            ListResponse = await GetStuff(baseUrl)
            $ListItems = ListResponse.data
            return 
        }

        let status = event.detail.replaceAll(" ", "_")
        $animeCurrentTab = status
        ListResponse = await GetStuff(baseUrl + "&status=" + status)
        $ListItems = ListResponse.data
        canRefresh = true

        console.log(baseUrl)
    }

    /**
     * @param event {Event}
     */
    async function handleScroll() {
        throttle(async () => {     
            const scrollTop = Math.floor(itemContainer.scrollTop) + 1000
            const scrollBottom = (itemContainer.scrollHeight - itemContainer.offsetHeight)

            if(scrollTop >= scrollBottom && canRefresh) {
                canRefresh = false

                if(ListResponse.paging.next === null || ListResponse.paging.next === undefined) return

                ListResponse = await GetStuff(ListResponse.paging.next)
                $ListItems = [...$ListItems, ...ListResponse.data]
                canRefresh = true
            }
        }, 100)
    }

    function closePopover() {
        return (event) => {
            if(event.target === $popover.element) {
                $popover.element.classList.add("inactive")
            }
        }
    }

    async function GetStuff(url) {
        const response = await GetRequest(url)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: response.data,
            });
            return
        }

        return JSON.parse(response.data)
    }

    var throttleTimer;
    const throttle = (callback, time) => {
        if (throttleTimer) return;
        throttleTimer = true;

        setTimeout(() => {
            callback();
            throttleTimer = false;
        }, time);
    };
</script>

<style lang="less">

    ::-webkit-scrollbar {
        width: 6px;
    }
    
    ::-webkit-scrollbar-thumb {
        background: #888; 
        border-radius: 1rem;
    }

    ::-webkit-scrollbar-thumb:hover {
        background: #555; 
    }

    main {
        height: calc(100% - 2.5rem);
        padding: 1rem;
        overflow-y: scroll;
        gap: 1rem;
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        align-content: flex-start;

        margin-right: 3px;
        transform: translate3d(0,0,0);
    }

    .popover {
        position: absolute;
        display: block;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        opacity: 1;
        user-select: none;
        pointer-events: all;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: rgba(0, 0, 0, 0.8);


        transition: all 0.2s ease-in-out;

        &.inactive {
            opacity: 0;
            pointer-events: none;
            user-select: none;
            background-color: rgba(0, 0, 0, 0);
            transition: all 0.2s ease-in-out;
        }

        .popover-content {
            background-color: var(--text-white);
            color: var(--text-black);
            padding: 1rem;
            border-radius: 0.5rem;
            width: 60%;
            max-width: 500px;

            h3 {
                font-size: 1.5rem;
                margin: 0;
                text-align: center;
                padding-bottom: 1rem;
            }

            p {
                margin: 0;
            }

            .options {
                display: grid;
                grid-template-columns: auto 1fr;
                gap: 1rem;

                input[type="number"] {
                    width: 4rem;
                }

                .inline {
                    display: flex;
                    align-items: center;
                    gap: 0.5rem;
                }
            }

            .buttons {
                display: flex;
                justify-content: end;
                padding-top: 1rem;
                gap: 0.5rem;

                button {
                    padding: 0.5rem 1rem;
                    border-radius: 0.5rem;
                    border: none;
                    background-color: var(--mal-blue);
                    color: var(--text-white);
                    cursor: pointer;
                    transition: all 0.2s ease-in-out;

                    &:hover {
                        background-color: var(--mal-blue-dark);
                    }
                }
            }
        }
    }

</style>
