<nav>
    <button on:click={() => {window.history.back()}}>
        <span class="material-symbols-outlined">
            arrow_back
        </span>
    </button>
</nav>

<main>
    <div class="flex">

        <div class="title block">
            <h1>{checkNull(animeData?.title)}</h1>
            <h2>{checkNull(animeData?.alternative_titles?.en)}</h2>
        </div>

        <div class="row main-info">

            <div class="fullwidth content">

                <div class="block score">
                    <p>Score</p>
                    <h2>{checkNull(animeData?.mean)}</h2>
                    <span>{checkNull(animeData?.num_scoring_users)} Users</span>
                </div>
                
                <div class="block rank">
                    <div class="top">
                        <p>Ranked <span class="bold">#{checkNull(animeData?.rank)}</span></p>
                        <p>Popularity <span class="bold">#{animeData?.popularity}</span></p>
                        <p>Members <span class="bold">
                            {formatNumber(
                                checkNull(animeData?.num_list_users)
                            )}
                        </span></p>
                    </div>
                    <div class="bottom">
                        {#if animeData?.start_season != null}
                            <p>{capitalize(checkNull(animeData?.start_season?.season))} {checkNull(animeData?.start_season?.year)}</p>
                            <hr>
                        {/if}

                        <p>{capitalize(checkNull(animeData?.media_type))}</p>
                        <hr>

                        <div class="studios">
                            {#if checkNull(animeData) !== "N/A"}
                                {#each animeData?.studios as studio, index}
                                    <p>{studio.name}</p>
                                    {#if index < animeData?.studios.length - 1}
                                        <span>&bull;</span>
                                    {/if}
                                {/each}
                            {/if}
                        </div>
                    </div>
                </div>

                <div class="block actions">

                    {#if checkNull(animeData) !== "N/A"}

                        {#if animeData?.my_list_status == null}
                            <button class="primary" on:click={handleAddToList} data-ani-id={animeData?.id}>Add to List</button>
                        {:else}
                            <div class="select-container">
                                <select id={`sel-status-${animeData?.id}`} on:change={handleStatusSelectChange} data-ani-id={animeData?.id}>
                                    {#each selectOptions as option (option)}
                                        <option value={option} selected={option === animeData?.my_list_status.status}>
                                            {capitalize(option)}
                                        </option>
                                    {/each}
                                </select>
                                <label for={`sel-status-${animeData?.id}`}>
                                    <span class="material-symbols-outlined">
                                        arrow_drop_down
                                    </span>
                                </label>
                            </div>
                        {/if}

                        <div class="select-container">
                            <select id={`sel-score-${animeData?.id}`} on:change={handleScoreSelectChange} data-ani-id={animeData?.id}>
                                {#each selectScores as score}
                                    <option value={score.value} selected={score.value === animeData?.my_list_status?.score}>
                                        {score.text}
                                    </option>
                                {/each}
                            </select>
                            <label for={`sel-score-${animeData?.id}`}>
                                <span class="material-symbols-outlined">
                                    arrow_drop_down
                                </span>
                            </label>
                        </div>

                        <div class="episodes">
                            <p>Episodes: </p>
                            <input type="number" on:focusout={handleEpisodesInputFocus} value={animeData.my_list_status ? animeData.my_list_status.num_episodes_watched : 0} data-ani-id={animeData?.id}>
                            <p> / {animeData.num_episodes}</p>
                            <button on:click={handleEpisodesButtonClick} data-ani-id={animeData?.id}>
                                <span class="material-symbols-outlined">add</span>
                            </button>
                        </div>

                    {/if}

                </div>

                <div class="block synopsis">
                    <h2>Synopsis</h2>
                    <p>{checkNull(animeData?.synopsis)}</p>
                </div>

            </div>

            <img src={animeData !== undefined ? animeData.main_picture.large : "/anime-placeholder.jpg"} alt={"poster " + checkNull(animeData?.title)}>

        </div>

        <div class="row extra-info">
            
            <div class="left block">

            </div>

            <div class="right fullwidth">

                <div class="photos block">
                    <!-- svelte-ignore a11y-no-static-element-interactions -->
                    {#if checkNull(animeData) !== "N/A"}
                        <div class="content">
                            {#each animeData?.pictures as picture}
                                <!-- svelte-ignore a11y-click-events-have-key-events -->
                                <!-- svelte-ignore a11y-no-noninteractive-element-to-interactive-role -->
                                <img src={picture.large} alt={picture.large} on:click={openPhotoPopup} role="button" tabindex="0">
                            {/each}
                        </div>
                    {/if}
                </div>

            </div>

        </div>
    </div>
</main>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="photo-popup hidden" bind:this={photoPopup} on:click={closePhotoPopup} role="button" tabindex="0">
    <div class="content">
        <img src="" alt="">
    </div>
</div>

<script>
    import { page } from "$app/stores";
    import { title } from "$lib/store";
    import { GetRequest, PatchRequest } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";
    import { Toast, capitalize, formatNumber, checkNull } from "$lib";

    let animeData = undefined;
    const fields = "status,genres,media_type,mean,rank,num_episodes,rating,my_list_status,alternative_titles,num_scoring_users,num_list_users,popularity,start_season,studios,synopsis,pictures"

    const selectOptions = ["watching", "completed", "on_hold", "dropped", "plan_to_watch"]
    const selectScores = [
        {value: 0, text: "Select"},
        {value: 10, text: "10 Masterpiece"},
        {value: 9, text: "9 Great"},
        {value: 8, text: "8 Very Good"},
        {value: 7, text: "7 Good"},
        {value: 6, text: "6 Fine"},
        {value: 5, text: "5 Average"},
        {value: 4, text: "4 Bad"},
        {value: 3, text: "3 Very Bad"},
        {value: 2, text: "2 Horrible"},
        {value: 1, text: "1 Appalling"}
    ]

    let photoPopup

    onMount(async () => {
        title.set("Anime Details");

        const url = $page.url
        const id = url.searchParams.get('id')

        let response = await GetRequest(`https://api.myanimelist.net/v2/anime/${id}?fields=${fields}`)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Could not find anime"
            })
            return
        }

        animeData = JSON.parse(response.data);
    })

    async function handleStatusSelectChange(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "status": event.target.value
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        handleResponse(response)
    }

    async function handleAddToList(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "status": "plan_to_watch"
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        handleResponse(response)
    }

    async function handleScoreSelectChange(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "score": event.target.value
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        handleResponse(response)
    }

    async function handleEpisodesButtonClick(event) {
        const id = event.target.attributes["data-ani-id"].value

        let episodesWatched = animeData.my_list_status ? animeData.my_list_status.num_episodes_watched : 0

        episodesWatched = episodesWatched + 1

        if(episodesWatched > animeData.num_episodes) {
            episodesWatched = animeData.num_episodes
        }

        const postData = {
            num_watched_episodes: episodesWatched
        }

        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        handleResponse(response)
    }

    async function handleEpisodesInputFocus(event) {
        const id = event.target.attributes["data-ani-id"].value
        let episodesWatched = event.target.value

        if(episodesWatched > animeData.num_episodes) {
            episodesWatched = animeData.num_episodes
        }

        const postData = {
            num_watched_episodes: episodesWatched
        }

        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        handleResponse(response)
    }

    /** @param {{success:boolean, data:any}} response */
    function handleResponse(response) {
        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Error updating list"
            })
            return
        }

        const data = JSON.parse(response.data)
        animeData.my_list_status = data

        Toast.fire({
            icon: "success",
            title: "Updated list"
        })
    }

    function openPhotoPopup(event) {
        const img = event.target
        photoPopup.querySelector("img").src = img.src
        photoPopup.classList.remove("hidden")
    }

    function closePhotoPopup() {
        photoPopup.classList.add("hidden")
    }
</script>

<style lang="less">

    nav {
        width: 100%;
        padding: 0.25rem;
        display: flex;
        height: 2.5rem;

        button {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 4rem;
            height: 100%;
            border-radius: 1rem;
            background: var(--mal-blue);
            color: white;
            border: none;
            cursor: pointer;
            transition: all 0.2s ease-in-out;
        }
    }

    main {
        height: calc(100% - 2.5rem);
        display: flex;
        justify-content: center;
        overflow-y: scroll;

        .flex {
            --_gap: 1rem;

            display: flex;
            flex-direction: column;
            gap: var(--_gap);
            padding: 1rem;
            height: 100%;
            width: 100%;
            max-width: 80rem;

            .row {
                display: flex;
                gap: var(--_gap);
            }

            .row:last-child {
                padding-bottom: 2.5rem;
            }

            .block {
                padding: 0.5rem 1rem;
                background: color-mix(in srgb, var(--text-white), var(--mal-blue) 10%);;
                border-radius: 0.5rem;
            }

            .fullwidth {
                flex-grow: 1;
            }

            span.bold {
                font-weight: 500;
            }

            .title {
                grid-area: "title";
                display: flex;
                flex-direction: column;

                h1 {
                    font-size: 2rem;
                    margin: 0;
                }

                h2 {
                    font-size: 1.5rem;
                    margin: 0;
                }
            }

            .main-info {

                img {
                    grid-area: "1-image";
                    width: 21rem;
                    height: 30rem;
                    min-width: 21rem;
                    min-height: 30rem;
                    object-fit: cover;
                    user-select: none;
                    border-radius: 0.25rem;
                    -webkit-user-drag: none;
                }

                .content {
                    height: 30rem;
                    display: grid;
                    gap: var(--_gap);
                    grid-template-areas: 
                    "score rank"
                    "actions actions"
                    "synopsis synopsis";

                    grid-template-rows: 8rem 3rem 17rem;
                    grid-template-columns: 8rem 1fr;

                    .score {
                        grid-area: score;
                        display: flex;
                        flex-direction: column;
                        justify-content: center;
                        align-items: center;

                        p {
                            margin: 0;
                            font-size: 1rem;
                            padding: 0.25rem 1rem;
                            background-color: var(--mal-blue);
                            color: var(--text-white);
                            border-radius: 0.5rem;
                        }

                        h2 {
                            margin: 0;
                            font-size: 2rem;
                        }
                    }

                    .rank {
                        grid-area: rank;
                        display: flex;
                        flex-direction: column;
                        justify-content: center;
                        gap: 1rem;

                        .top {
                            display: flex;
                            gap: 1.5rem;

                            p {
                                margin: 0;
                                font-size: 1.2rem;
                            }
                        }

                        .bottom {
                            display: flex;
                            align-items: center;
                            gap: 1rem;

                            height: 1.5rem;
                        
                            hr {
                                height: 100%;
                                margin: 0.5rem 0;
                                border: none;
                                border-left: 2px solid #c6c6c6;
                            }

                            p {
                                margin: 0;
                                font-size: 1rem;
                                font-weight: 500;
                                color: var(--mal-blue);
                            }

                            .studios {
                                display: flex;
                                flex-wrap: wrap;
                                gap: 0.5rem;
                            }
                        }
                    }

                    .actions {
                        --_button-height: 2rem;
                        --_border_radius: 0.5rem;
                        grid-area: actions;
                        display: flex;
                        align-items: center;
                        gap: 1rem;

                        .select-container {
                            display: flex;
                            --_border: 2px solid var(--mal-blue);

                            select {
                                all: unset;
                                padding: 0 0rem 0 1rem;
                                border-top-left-radius: 0.5rem;
                                border-bottom-left-radius: 0.5rem;
                                color: var(--text-black);
                                cursor: pointer;
                                transition: all 0.2s ease-in-out;
                                display: flex;
                                align-items: center;
                                border: var(--_border);
                                border-right: none;
                            }

                            label {
                                height: var(--_button-height);
                                border-top-right-radius: 0.5rem;
                                border-bottom-right-radius: 0.5rem;
                                color: var(--text-black);
                                cursor: pointer;
                                transition: all 0.2s ease-in-out;
                                display: flex;
                                align-items: center;
                                border: var(--_border);
                                border-left: none;

                                span {
                                    user-select: none;
                                    pointer-events: all;
                                }
                            }
                        }

                        button {
                            all: unset;
                            height: var(--_button-height);
                            padding: 0 1rem;
                            border-radius: 0.5rem;
                            border: none;
                            background-color: var(--mal-blue);
                            color: var(--text-white);
                            cursor: pointer;
                            transition: all 0.2s ease-in-out;

                            display: flex;
                            justify-content: center;
                            align-items: center;

                            &:hover {
                                background-color: var(--mal-blue-dark);
                            }
                        }

                        .episodes {
                            display: flex;
                            align-items: center;
                            border: solid 2px var(--mal-blue);
                            border-radius: 5px;
                            height: 2rem;
                            padding-inline: 0.5rem;

                            p {
                                margin: 0;
                            }

                            input {
                                border: none;
                                outline: none;
                                background: transparent;
                                width: 50px;
                                text-align: right;
                                font-size: 1rem;
                            }

                            input[type="number"]::-webkit-inner-spin-button,
                            input[type="number"]::-webkit-outer-spin-button {
                                -webkit-appearance: none;
                                margin: 0;
                            }

                            button {
                                all: unset;
                                height: 100%;
                                display: flex;
                                align-items: center;
                                cursor: pointer;

                                span {
                                    pointer-events: none;
                                }
                            }
                        }
                    }

                    .synopsis {
                        grid-area: synopsis;
                        display: flex;
                        flex-direction: column;
                        gap: 0.5rem;
                        height: 100%;
                        position: relative;

                        h2 {
                            margin: 0;
                            font-size: 1.5rem;
                        }

                        p {
                            --_scrollbar-width: 0.5rem;
                            margin: 0;
                            white-space: pre-wrap;
                            padding-bottom: 2rem;

                            overflow-y: auto;

                            &::-webkit-scrollbar {
                                width: var(--_scrollbar-width);
                            }

                            &::-webkit-scrollbar-thumb {
                                background-color: gray;
                                border-radius: 0.25rem;
                            }
                        }
                    }
                }
            }

            .extra-info {
                --_left-width: 21rem;

                .left {
                    width: 21rem;
                }

                .right {
                    width: calc(100% - var(--_left-width));
                    display: flex; 
                    flex-direction: column;
                    gap: var(--_gap);
                    padding: 0;

                    .photos {
                        .content {
                            display: flex;
                            flex-wrap: nowrap;
                            gap: 1rem;
                            overflow-x: auto;
                            user-select: none;

                            img {
                                height: 15rem;
                                object-fit: cover;
                                border-radius: 0.25rem;
                                margin-bottom: 5px;
                                -webkit-user-drag: none;
                                cursor: pointer;
                            }

                            &::-webkit-scrollbar {
                                width: 1px;
                                height: 5px;
                            }

                            &::-webkit-scrollbar-thumb {
                                background-color: gray;
                                border-radius: 0.25rem;
                            }
                        }                        
                    }
                }

            }
        }
    }

    .photo-popup {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;

        transition: all 0.2s ease-in-out;
        cursor: pointer;

        .content {
            width: 80%;
            height: 80%;
            display: flex;
            justify-content: center;
            align-items: center;

            img {
                max-width: 100%;
                max-height: 100%;
                object-fit: contain;
                -webkit-user-drag: none;
                user-select: none;
                border-radius: 0.5rem;
            }
        }

        &.hidden {
            opacity: 0;
            pointer-events: none;
        }
    }

</style>