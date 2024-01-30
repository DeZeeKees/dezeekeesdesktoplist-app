<div class="item">

    <div class="left">
        <img src={data.node.main_picture.medium} alt={data.node.title} draggable="false" loading="lazy">
    </div>

    <div class="right">
        <h2>#{data.node.rank} {data.node.title}</h2>
        <div>{@html formatGenres(data.node.genres)}</div>
        <p>
            {@html capitalize(data.node.media_type)} -
            {data.node.mean == null ? "N/A" : data.node.mean.toFixed(2)} -
            {@html formatStatus(data.node.status)}
        </p>
        <p>Total Episodes: {data.node.num_episodes > 0 ? data.node.num_episodes : "?"}</p>
        <p>Rating: {formatRating(data.node.rating)}</p>
        <div class="buttons">
            {#if data.node.my_list_status == null}
                <button class="primary" on:click={handleAddToList} data-ani-id={data.node.id}>Add to List</button>
            {:else}
            <div class="select-container">
                <select id={`sel-${data.node.id}`} on:change={handleSelectChange} data-ani-id={data.node.id}>
                    {#each selectOptions as option (option)}
                        <option value={option} selected={option === data.node.my_list_status.status}>
                            {capitalize(option)}
                        </option>
                    {/each}
                </select>
                <label for={`sel-${data.node.id}`}>
                    <span class="material-symbols-outlined">
                        arrow_drop_down
                    </span>
                </label>
            </div>
            {/if}
            <a href="#none">View Anime</a>
        </div>
    </div>

</div>

<script>
    import { Toast } from "$lib";
    import { PatchRequest } from "$lib/wailsjs/go/main/App";

    export let data = {}

    let selectOptions = ["watching", "completed", "on_hold", "dropped", "plan_to_watch"]

    function formatStatus(status) {
        switch(status) {
            case "finished_airing":
                return "<span style='color: #3F51B5;'>Finished Airing</span>";
            case "currently_airing":
                return "<span style='color: #4CAF50;'>Currently Airing</span>";
            case "not_yet_aired":
                return "<span style='color: #F44336;'>Not Yet Aired</span>";
            default:
                return status;
        }
    }

    function formatGenres(genres) {
        let genresString = ""

        genres.map(genre => {
            genresString += "<span class='chip'>" + genre.name + "</span>";
        })

        return genresString;
    }

    function capitalize(string) {
        // replace all underscores with spaces
        string = string.replace(/_/g, " ");
        // capitalize first letter
        return string.charAt(0).toUpperCase() + string.slice(1);
    }

    function formatRating(rating) {
        switch(rating) {
            case "g":
                return "G - All Ages";
            case "pg":
                return "PG - Children";
            case "pg_13":
                return "PG-13 - Teens 13 or older";
            case "r":
                return "R - 17+ (violence & profanity)";
            case "r+":
                return "R+ - Mild Nudity";
            case "rx":
                return "Rx - Hentai";
            default:
                return "N/A";
        }
    }

    async function handleSelectChange(event) {
        const id = event.target.attributes["data-ani-id"].value
        const postData = {
            "status": event.target.value
        }
        const url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        const result = await PatchRequest(url, JSON.stringify(postData)).then((data) => {
            if(data === "error") {
                Toast.fire({
                    icon: "error",
                    title: "Error updating list"
                })
                return null
            }

            return JSON.parse(data)
        })

        if(result === null) return

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

        const result = await PatchRequest(url, JSON.stringify(postData)).then((data) => {
            if(data === "error") {
                Toast.fire({
                    icon: "error",
                    title: "Error adding to list"
                })
                return null
            }

            return JSON.parse(data)
        })

        if(result === null) return

        Toast.fire({
            icon: "success",
            title: "Added to list"
        })

        data.node.my_list_status = {}
        data.node.my_list_status.status = "plan_to_watch"
    }
</script>

<style lang="less">

    .item {
        --_border-radius: 0.5rem;
        height: 250px;
        width: 100%;
        display: flex;
        border-radius: var(--_border-radius);
        background: color-mix(in srgb, var(--text-white), var(--mal-blue) 10%);

        .left {
            height: 250px; 
            aspect-ratio: 5/7;
            border-radius: var(--_border-radius) 0 0 var(--_border-radius);
            overflow: hidden;
            flex-shrink: 0;

            img {
                height: 100%;
                width: 100%;
                object-fit: cover;
            }
        }

        .right {
            flex-grow: 1;
            overflow: hidden;
            padding: 0.5rem 1rem;
            display: flex;
            flex-direction: column;

            h2 {
                margin: 0;
                font-size: 1.5rem;
                font-weight: 500;
                overflow-x: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }

            p {
                margin: 0;
                font-size: 1rem;
                font-weight: 400;
                overflow-x: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }

            :global(span.chip) {
                display: inline-block;
                padding: 0.25rem 0.5rem;
                margin: 0.25rem 0.25rem 0.25rem 0;
                border-radius: 0.5rem;
                background-color: var(--mal-blue);
                color: var(--text-white);
                font-size: 0.75rem;
                font-weight: 500;
            }

            .buttons {
                --_button-height: 2rem;
                flex-grow: 1;
                display: flex;
                flex-direction: row;
                justify-content: flex-end;
                align-items: flex-end;
                gap: 0.5rem;

                button, a {
                    all: unset;
                    height: var(--_button-height);
                    padding: 0 1rem;
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
            }
        }
    }

</style>