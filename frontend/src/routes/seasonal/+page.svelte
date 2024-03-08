<Tabs renderTabs={false} />

<main>
    {#if seasonalData === undefined}
        <div class="center">
            <h1>Loading...</h1>
        </div>
    {:else}
        
        <div class="seasonal-container">

            <div class="seasonal-items">
                <h1>TV (new)</h1>
                {#if currentSeasonalItems.length === 0}
                    <p>No current season items found</p>
                {:else}
                    {#each currentSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <div class="seasonal-items">
                <h1>TV (continuing)</h1>
                {#if continuingSeasonalItems.length === 0}
                    <p>No continuing season items found</p>
                {:else}
                    {#each continuingSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <div class="seasonal-items">
                <h1>ONA's</h1>
                {#if onaSeasonalItems.length === 0}
                    <p>No ona items found</p>
                {:else}
                    {#each onaSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <div class="seasonal-items">
                <h1>OVA's</h1>
                {#if ovaSeasonalItems.length === 0}
                    <p>No ova items found</p>
                {:else}
                    {#each ovaSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <div class="seasonal-items">
                <h1>Movies</h1>
                {#if movieSeasonalItems.length === 0}
                    <p>No movie items found</p>
                {:else}
                    {#each movieSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <div class="seasonal-items">
                <h1>Specials</h1>
                {#if specialSeasonalItems.length === 0}
                    <p>No special items found</p>
                {:else}
                    {#each specialSeasonalItems as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div>

            <!-- <div class="seasonal-items">
                <h1>Other</h1>
                {#if seasonalData.data.length === 0}
                    <p>No special items found</p>
                {:else}
                    {#each seasonalData.data as item}
                        <img src={item.node.main_picture.medium} alt="">
                    {/each}
                {/if}
            </div> -->

        </div>

    {/if}
</main>

<script>
    import Tabs from "$lib/components/tabs.svelte";
    import { afterNavigate } from "$app/navigation";
    import { title } from "$lib/store";
    import { Toast, capitalize } from "$lib";
    import { page } from "$app/stores";
    import { GetRequest, GetSettings } from "$lib/wailsjs/go/main/App";

    const seasonMap = {
        1: "winter",
        2: "winter",
        3: "winter",
        4: "spring",
        5: "spring",
        6: "spring",
        7: "summer",
        8: "summer",
        9: "summer",
        10: "fall",
        11: "fall",
        12: "fall"
    }

    const CurrentMonth = new Date().getMonth() + 1
    const CurrentYear = new Date().getFullYear()

    const fields = "start_season,media_type"

    let seasonalData = undefined;

    let currentSeasonalItems = []
    let continuingSeasonalItems = []
    let onaSeasonalItems = []
    let ovaSeasonalItems = []
    let movieSeasonalItems = []
    let specialSeasonalItems = []

    afterNavigate(async () => {
        const url = $page.url
        const season = url.searchParams.get('season') || seasonMap[CurrentMonth]
        const year = url.searchParams.get('year') || CurrentYear

        const settings = await GetSettings()

        title.set(`Seasonal - ${capitalize(season)} ${year}`)

        const FetchUrl = `https://api.myanimelist.net/v2/anime/season/${year}/${season}?fields=${fields}&limit=500&sort=anime_num_list_users&nsfw=${settings.nsfwContent}`

        console.log(FetchUrl)

        const response = await GetRequest(FetchUrl)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: `Something went wrong fetching ${season} ${year} seasonal data. Please try again later.`
            })
            return
        }

        seasonalData = JSON.parse(response.data);
        
        for(const item of seasonalData.data) {

            const { node } = item
            
            // current season tv items
            if(node.start_season?.year === year && node.start_season?.season === season && node.media_type === "tv") {
                currentSeasonalItems.push(item)
            }

            // continuous season tv items
            if(node.start_season?.year !== year && node.media_type === "tv") {
                continuingSeasonalItems.push(item)
            }

            // ona items
            if(node.media_type === "ona") {
                onaSeasonalItems.push(item)
            }

            // ova items
            if(node.media_type === "ova") {
                ovaSeasonalItems.push(item)
            }

            // movie items
            if(node.media_type === "movie") {
                movieSeasonalItems.push(item)
            }

            // special items
            if(node.media_type === "special") {
                specialSeasonalItems.push(item)
            }

        }
        
    })

</script>

<style lang="less">

    main {
        height: calc(100% - 2.5rem);
        display: flex;
        justify-content: center;
        overflow-y: scroll;
    }

</style>