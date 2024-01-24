<Tabs tabs={tabList} on:tabChange={handleTabChange} />

<main>
    {#if currentTab === "general" && releaseInfo !== undefined}
        <h2>General</h2>
        <p>Use Prerelease</p>
        <input type="checkbox">
    {:else if currentTab === "updates" && releaseInfo !== undefined}
        <div class="updates">
            
            {#if releaseInfo.isLatest}
                <h2>You are up to date</h2>
                <p>Current version: {version}</p>
                <p>Release date: {formateStringDate(releaseInfo.release.published_at)}</p>
                <div class="buttons">
                    <button on:click={() => BrowserOpenURL(releaseInfo.release.html_url)}>Release notes</button>
                </div>
            {:else}
                <h2>Update available</h2>
                <p>Current version: {version}</p>
                <p>Latest version: {releaseInfo.release.tag_name}</p>
                <p>Release date: {formateStringDate(releaseInfo.release.published_at)}</p>
                <div class="buttons">
                    <button on:click={() => BrowserOpenURL(releaseInfo.release.html_url)}>Release notes</button>
                    <button on:click={handleDownload}>Download</button>
                </div>
            {/if}

        </div>
    {:else}
        <h1>Loading...</h1>
    {/if}
</main>

<script>
    import { GetVersion, GetReleaseInfo, InstallUpdate } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import Tabs from "$lib/components/tabs.svelte";
    import { title } from "$lib/store";
    import { BrowserOpenURL } from "$lib/wailsjs/runtime/runtime";

    let version;
    let releaseInfo = undefined;
    let currentTab = "general";

    const tabList = [
        {name: "General"},
        {name: "Updates"},
    ]

    onMount(async () => {
        version = await GetVersion();
        releaseInfo = await GetReleaseInfo(true);

        console.log(releaseInfo);

        title.set(`Settings - General`);
    });

    function handleTabChange(event) {
        currentTab = event.detail;
        title.set(`Settings - ${capitalizeFirstLetter(event.detail)}`);
    }

    function capitalizeFirstLetter(string) {
        return string.charAt(0).toUpperCase() + string.slice(1);
    }

    function formateStringDate(date) {
        date = new Date(date);
        return `${date.getDate()}-${(date.getMonth() + 1)}-${date.getFullYear()}`
    }

    async function handleDownload() {
        const data = await InstallUpdate();

        console.log(data);
    }
</script>

<style lang="less">

main {
    display: flex;
    flex-direction: column;
    align-items: center;
    height: calc(100vh - 2.5rem);
    padding-top: 4rem;

    .updates {
        background-color: #fff; /* Set a background color for the card */
        border-radius: 8px; /* Add border-radius for rounded corners */
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3); /* Add a subtle box shadow for depth */
        padding: 20px; /* Add padding inside the card */
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        max-width: 600px;

        h2,
        p {
            color: var(--text-black);
            margin: 0;
        }

        .buttons {
            display: flex;
            flex-direction: row;
            justify-content: flex-end;
            margin-top: 20px;
        }

        button {
            background-color: var(--mal-blue);
            color: var(--text-white);
            border: none;
            padding: 12px 24px;
            cursor: pointer;
            font-size: 16px;
            font-weight: bold;
            text-transform: uppercase;
            border-radius: 4px;
            transition: background-color 0.3s ease;
            margin-right: 10px;

            &:hover {
                background-color: var(--mal-blue-dark);
            }
        }

        button:nth-child(2) {
            background-color: var(--mal-blue);
            color: var(--text-white);
            border: none;
            padding: 12px 24px;
            cursor: pointer;
            font-size: 16px;
            font-weight: bold;
            text-transform: uppercase;
            border-radius: 4px;
            transition: background-color 0.3s ease;

            &:hover {
                background-color: var(--mal-blue-dark);
            }
        }
    }
}

</style>