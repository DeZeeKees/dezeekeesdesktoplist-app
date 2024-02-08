<Tabs renderTabs={false}>
    <div class="search-container">
        <div class="search-bar">
            <input type="search" name="anime_search" id="anime_search" placeholder="Search Anime" bind:this={searchInput} on:keyup={handleSearch}>
            <button class="search-button" on:click={handleSearchClick}>
                <span class="material-symbols-outlined">
                    search
                </span>
            </button>
        </div>
    </div>
</Tabs>

<main>
    {#if searchData === undefined}
        <div class="center">
            <h1>Search for an anime</h1>
        </div>
    {:else if searchItems.length === 0}
        <div class="center">
            <h1>No results found</h1>
        </div>
    {:else}
        {#each searchItems as searchItem}
            <SearchItem data={searchItem}/>
        {/each}
    {/if}
</main>

<script>
    import Tabs from "$lib/components/tabs.svelte";
    import SearchItem from "$lib/components/SearchItem.svelte";
    import { title } from "$lib/store";
    import { GetRequest, GetSettings } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import { Toast } from "$lib";

    let searchData = undefined;
    let searchItems = [];
    /** @type {HTMLInputElement} */
    let searchInput

    const fields = "status,genres,media_type,mean,rank,num_episodes,rating,my_list_status"

    let settings = {
        usePrerelease: false,
        yourListCardSizeMultiplier: 1,
        nsfwContent: false,
    }

    onMount(async () => {
        title.set(`Search`);

        settings = await GetSettings()
    })

    async function handleSearch(event) {
        let searchValue = searchInput.value.replace(/\s/g, "");

        // if both are false pass the if statement
        if (event.key !== "Enter" || searchValue.length < 3) {
            return;
        }

        const response = await GetRequest("https://api.myanimelist.net/v2/anime?limit=30&fields=" + fields +"&q=" + searchValue + "&nsfw=" + settings.nsfwContent)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Something went wrong"
            })
            return
        }

        searchData = JSON.parse(response.data);
        
        searchItems = searchData.data;
    }

    async function handleSearchClick() {
        let searchValue = searchInput.value.replace(/\s/g, "");

        if(searchValue.length < 3) {
            return;
        }

        const response = await GetRequest("https://api.myanimelist.net/v2/anime?limit=30&fields=" + fields +"&q=" + searchValue + "&nsfw=" + settings.nsfwContent)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Something went wrong"
            })
            return
        }

        searchData = JSON.parse(response.data);
        
        searchItems = searchData.data;
    }
</script>

<style lang="less">

    .search-container {
        display: flex;
        flex-direction: column;
        align-items: end;
        justify-content: center;
        height: 100%;
    }

    .search-bar {
        --_border-radius: 10rem;
        width: 100%;
        max-width: 500px;
        height: 100%;
        font-size: 1rem;
        outline: none;
        display: flex;
        box-shadow: 0 0 0.5rem rgba(0, 0, 0, 0.3);
        border-radius: var(--_border-radius);

        button {
            width: 3rem;
            height: 100%;
            border: none;
            border-radius: 0 var(--_border-radius) var(--_border-radius) 0;
            background-color: var(--mal-blue);
            color: var(--text-white);
            cursor: pointer;
            transition: background-color 0.3s ease;

            &:hover {
                background-color: var(--mal-blue-dark);
            }

            span {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 100%;
                height: 100%;
            }
        }

        input[type="search"] {
            width: calc(100% - 3rem);
            height: 100%;
            padding: 1rem;
            border: none;
            border-radius: var(--_border-radius) 0 0 var(--_border-radius);
            outline: none;
            font-size: 1rem;
        }

        input[type="search"]::placeholder {
            color: var(--text-dark);
        }
    }

    main {
        height: calc(100% - 2.5rem);
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        padding: 1rem;
        overflow-y: scroll;
        
        margin-right: 3px;
        transform: translate3d(0,0,0);
    }

    .center {
        width: 100%;
        display: flex;
        justify-content: center;
    }

</style>