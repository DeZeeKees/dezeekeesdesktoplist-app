<Tabs renderTabs={false}>
    <div class="search-container">
        <div class="search-bar">
            <input type="search" name="anime_search" id="anime_search" placeholder="Search Anime" bind:this={searchInput} on:input={handleSearch}>
            <button class="search-button">
                <span class="material-symbols-outlined">
                    search
                </span>
            </button>
        </div>
    </div>
</Tabs>

<main>
    {#if searchData === undefined}
        <div>
            <h1>Search for an anime</h1>
        </div>
    {:else if searchData.data.length === 0}
        <div>
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
    import { GetRequest } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";

    let searchData = undefined;
    let searchItems = [];
    /** @type {HTMLInputElement} */
    let searchInput

    onMount(async () => {
        title.set(`Search`);
    })

    let timeout = setTimeout(() => {}, 0);

    async function handleSearch() {
        clearTimeout(timeout);
        timeout = setTimeout(async () => {
            if (searchInput.value.length > 2) {
                searchData = await GetRequest("https://api.myanimelist.net/v2/anime?limit=10&q=" + searchInput.value)
                    .then(data => JSON.parse(data));

                searchItems = searchData.data;
            } else {
                searchData = undefined;
            }
        }, 500);
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

</style>