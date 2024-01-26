<Tabs tabs={tabList} on:tabChange={handleTabChange} />

<main>
    {#if currentTab === "general" && releaseInfo !== undefined}
        <div class="settings">
            <h2>General</h2>

            <div class="row">
                <p>Use Pre-release</p>
                <label class="toggle-switch" for="use_prerelease" >
                    <input type="checkbox" name="use_prerelease" id="use_prerelease" bind:checked={settings.usePrerelease}>
                    <span class="toggle-button round"></span>
                </label>
            </div>
            
            <h2>Your List</h2>
            
            <div class="row">
                <label for="card_size_multiplier">Card Size Multiplier {settings.yourListCardSizeMultiplier}</label>
                <div class="slider">
                    <button on:click={handleSliderReset}>Reset</button>
                    <input class="" type="range" min="40" max="200" name="card_size_multiplier" on:input={handleSliderChange} bind:this={CardSizeRangeElement}>
                </div>
            </div>

            <div class="buttons">
                <button on:click={handleSaveSettings}>Save</button>
            </div>
        </div>
    {:else if currentTab === "updates" && releaseInfo !== undefined}
        <div class="updates_container">
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
        </div>
    {:else}
        <h1>Loading...</h1>
    {/if}
</main>

<script>
    import { GetVersion, GetReleaseInfo, InstallUpdate, GetSettings, SaveSettings } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import Tabs from "$lib/components/tabs.svelte";
    import { title } from "$lib/store";
    import { BrowserOpenURL } from "$lib/wailsjs/runtime/runtime";
    import { Toast } from "$lib";

    let version;
    let releaseInfo = undefined;
    let currentTab = "general";

    const tabList = [
        {name: "General"},
        {name: "Updates"},
    ]

    let settings = {
        usePrerelease: false,
        yourListCardSizeMultiplier: 1
    }

    /** @type {HTMLInputElement}*/
    let CardSizeRangeElement;

    onMount(async () => {
        version = await GetVersion();
        releaseInfo = await GetReleaseInfo(true);
        settings = await GetSettings();

        CardSizeRangeElement.value = (settings.yourListCardSizeMultiplier * 100).toString();

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

        if(isNaN(date.getDate())) {
            return "Unknown";
        }

        return `${date.getDate()}-${(date.getMonth() + 1)}-${date.getFullYear()}`
    }

    async function handleDownload() {
        const data = await InstallUpdate();

        console.log(data);
    }

    async function handleSaveSettings() {
        const data = await SaveSettings(JSON.stringify(settings));

        if(data !== "success") {
            Toast.fire({
                icon: "error",
                title: "Failed to save settings",
            });
        }

        Toast.fire({
            icon: "success",
            title: "Settings saved",
        });
    }

    function handleSliderReset() {
        CardSizeRangeElement.value = "100";
        settings.yourListCardSizeMultiplier = 1;
    }

    function handleSliderChange(event) {
        settings.yourListCardSizeMultiplier = event.target.value / 100;
    }
</script>

<style lang="less">

.updates_container {
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

.settings {
    padding-inline: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;

    h2, p, label {
        margin: 0;
        padding-bottom: 0.1rem;
    }

    button {
        background-color: var(--mal-blue);
        color: var(--text-white);
        border: none;
        padding: 4px 8px;
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

    .row {
        display: flex;
        flex-direction: column;
    }

    .slider {
        display: flex;
        flex-direction: row;
        align-items: center;
        
        // type range
        input[type="range"] {
            -webkit-appearance: none;
            appearance: none;
            height: 20px;
            width: 100%;
            max-width: 800px;
            background-color: var(--text-light);
            outline: none;
            border-radius: 100px;
            opacity: 0.7;
            transition: opacity 0.2s ease-in-out;

            &:hover {
                opacity: 1;
            }

            &::-webkit-slider-thumb {
                appearance: none;
                width: 20px;
                height: 20px;
                background-color: var(--mal-blue);
                border-radius: 100%;
                cursor: pointer;
            }
        }
    }

    /* The switch - the box around the slider */
    .toggle-switch {
        position: relative;
        display: inline-block;
        width: 60px;
        height: 34px;
        cursor: pointer;
    }

    /* Hide default HTML checkbox */
    .toggle-switch input {
        opacity: 0;
        width: 0;
        height: 0;
    }

    /* The toggle button */
    .toggle-button {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: #ccc;
        transition: .4s;
        pointer-events: none;
    }

    .toggle-button:before {
        position: absolute;
        content: "";
        height: 26px;
        width: 26px;
        left: 4px;
        bottom: 4px;
        background-color: white;
        transition: .4s;
    }

    input:checked + .toggle-button {
        background-color: var(--mal-blue);
    }

    input:focus + .toggle-button {
        box-shadow: 0 0 1px var(--mal-blue);
    }

    input:checked + .toggle-button:before {
        transform: translateX(26px);
    }

    /* Rounded toggle button */
    .toggle-button.round {
        border-radius: 34px;
    }

    .toggle-button.round:before {
        border-radius: 50%;
    }
}

</style>