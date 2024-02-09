<nav>
    <button on:click={() => {window.history.back()}}>
        <span class="material-symbols-outlined">
            arrow_back
        </span>
    </button>
</nav>

<main>
    {#if animeData === undefined}
    <div class="flex">

        <div class="title block">
            <h1>Anime Title</h1>
            <h2>English Title</h2>
        </div>

        <div class="row">
            <img src="/anime-placeholder.jpg" alt="Empty anime poster">

            <div class="fullwidth content">

                <div class="block score">
                    <p>Score</p>
                    <h2>N/A</h2>
                    <span>N/A Users</span>
                </div>
                
                <div class="block rank">
                    <div class="top">
                        <p>Ranked <span class="bold">#N/A</span></p>
                        <p>Popularity <span class="bold">#N/A</span></p>
                        <p>Members <span class="bold">N/A</span></p>
                    </div>
                    <div class="bottom">

                    </div>
                </div>

                <div class="block actions">
                    <button class="primary">Add to List</button>
                </div>

                <div class="block synopsis">
                    <h2>Synopsis</h2>
                    <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Et facere id nisi vitae, necessitatibus molestias aliquid qui ipsum rerum soluta ad rem, amet cupiditate nostrum laborum laudantium enim modi sapiente?</p>
                </div>
            </div>

        </div>

    </div>
    {:else}
        <div class="flex">

            <div class="title block">
                <h1>{animeData?.title}</h1>
                <h2>{animeData?.alternative_titles.en}</h2>
            </div>

            <div class="row">
                <img src={animeData?.main_picture.large} alt={"poster " + animeData?.title}>

                <div class="fullwidth content">

                    <div class="block score">
                        <p>Score</p>
                        <h2>{animeData?.mean}</h2>
                        <span>{animeData?.num_scoring_users} Users</span>
                    </div>
                    
                    <div class="block rank">
                        <div class="top">
                            <p>Ranked <span class="bold">#{animeData?.rank}</span></p>
                            <p>Popularity <span class="bold">#{animeData?.popularity}</span></p>
                            <p>Members <span class="bold">{formatNumber(animeData?.num_list_users)}</span></p>
                        </div>
                        <div class="bottom">
                            {#if animeData?.start_season != null}
                                <p>{capitalize(animeData?.start_season.season)} {animeData?.start_season.year}</p>
                                <hr>
                            {/if}

                            <p>{capitalize(animeData?.media_type)}</p>
                            <hr>

                            <div class="studios">
                                {#each animeData?.studios as studio, index}
                                    <p>{studio.name}</p>
                                    {#if index < animeData?.studios.length - 1}
                                        <span>&bull;</span>
                                    {/if}
                                {/each}
                            </div>
                        </div>
                    </div>

                    <div class="block actions">

                        {#if animeData?.my_list_status == null}
                            <button class="primary" on:click={handleAddToList} data-ani-id={animeData?.id}>Add to List</button>
                        {:else}
                            <div class="select-container">
                                <select id={`sel-${animeData?.id}`} on:change={handleSelectChange} data-ani-id={animeData?.id}>
                                    {#each selectOptions as option (option)}
                                        <option value={option} selected={option === animeData?.my_list_status.status}>
                                            {capitalize(option)}
                                        </option>
                                    {/each}
                                </select>
                                <label for={`sel-${animeData?.id}`}>
                                    <span class="material-symbols-outlined">
                                        arrow_drop_down
                                    </span>
                                </label>
                            </div>
                        {/if}

                    </div>

                    <div class="block synopsis">
                        <h2>Synopsis</h2>
                        <p>{animeData?.synopsis}</p>
                    </div>

                </div>

            </div>
        </div>
    {/if}
</main>

<script>
    import { page } from "$app/stores";
    import { Toast } from "$lib";
    import { title } from "$lib/store";
    import { GetRequest, PatchRequest } from "$lib/wailsjs/go/main/App";
    import { onMount } from "svelte";

    let animeData = undefined;
    const fields = "status,genres,media_type,mean,rank,num_episodes,rating,my_list_status,alternative_titles,num_scoring_users,num_list_users,popularity,start_season,studios,synopsis"

    const selectOptions = ["watching", "completed", "on_hold", "dropped", "plan_to_watch"]

    onMount(async () => {
        title.set("Anime Details");

        let response = await GetRequest(`https://api.myanimelist.net/v2/anime/${$page.params.id}?fields=${fields}`)

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Could not find anime"
            })
            return
        }

        animeData = JSON.parse(response.data);

        console.log(animeData)
    })

    function formatNumber(num) {
        return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    }

    function capitalize(string) {
        // replace all underscores with spaces
        string = string.replace(/_/g, " ");
        // capitalize first letter
        return string.charAt(0).toUpperCase() + string.slice(1);
    }

    async function handleSelectChange(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "status": event.target.value
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        if(!response.data) {
            Toast.fire({
                icon: "error",
                title: "Error updating list"
            })
            return
        }

        Toast.fire({
            icon: "success",
            title: "Updated list"
        })
    }

    async function handleAddToList(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "status": "plan_to_watch"
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const response = await PatchRequest(url, JSON.stringify(postData))

        if(!response.success) {
            Toast.fire({
                icon: "error",
                title: "Error adding to list"
            })
            return
        }

        Toast.fire({
            icon: "success",
            title: "Added to list"
        })

        animeData.my_list_status = {}
        animeData.my_list_status.status = "plan_to_watch"
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

            img {
                grid-area: "1-image";
                width: 21rem;
                height: 30rem;
                object-fit: cover;
                user-select: none;
                border-radius: 0.25rem;
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
                }

                .synopsis {
                    grid-area: synopsis;
                    display: flex;
                    flex-direction: column;
                    gap: 0.5rem;

                    height: 100%;

                    h2 {
                        margin: 0;
                        font-size: 1.5rem;
                    }

                    p {
                        margin: 0;
                        overflow-y: scroll;

                        // keep space and new lines in text
                        white-space: pre-wrap;

                        &::-webkit-scrollbar {
                            width: 0.5rem;
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

    .center {
        display: flex;
        justify-content: center;
        align-items: center;
    }

</style>