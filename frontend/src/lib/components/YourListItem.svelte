<div class="item" bind:this={item} on:dblclick={handleEdit} role="button" tabindex="0">
    <img src={data.node.main_picture.medium} alt={data.node.title} draggable="false" loading="lazy">

    <div class="overlay">

        <div class="top">
            <div class="row one">
                <p class="type">
                    <span class="material-symbols-outlined">tv</span>
                    {data.node.media_type.charAt(0).toUpperCase() + data.node.media_type.slice(1)}
                </p>

                <button on:mouseup={handleEdit}>Edit</button>
            </div>

            <div class="row two">
                <p>Start: {data.node.start_date == undefined ? "?" : formatDate(data.node.start_date)}</p>
            </div>

            <div class="row three">
                <p>End: {data.node.end_date == undefined ? "?" : formatDate(data.node.end_date)}</p>
            </div>
        </div>

        <div class="bottom">
            <h3 class="title">{data.node.title}</h3>

            <div class="split">
                <div class="left">
                    <p class="score">{data.list_status.score == 0 ? "Not Scored" : `Score ${data.list_status.score}`}</p>
                </div>
                <div class="right">
                    <p class="eps">
                        {#if 
                            data.list_status.status == "watching" ||
                            data.list_status.status == "dropped" ||
                            data.list_status.status == "on_hold"
                        }
                        <div class="plus">
                            <div class="hover">
                                {data.list_status.num_episodes_watched} of {data.node.num_episodes == 0 ? "?" : data.node.num_episodes} <button on:mouseup={() => { incrementEpisodes(data.node.id) }}><span class="material-symbols-outlined">add</span></button>
                            </div>
                            <div class="nohover">
                                {data.list_status.num_episodes_watched} of {data.node.num_episodes == 0 ? "?" : data.node.num_episodes} eps
                            </div>
                        </div>
                        {:else}
                            {data.node.num_episodes == 0 ? "Unknown" : data.node.num_episodes + "eps"} 
                        {/if}
                    </p>
                </div>
            </div>
        </div>

    </div>
       
</div>

<script>
    import { PatchRequest } from "$lib/wailsjs/go/main/App";
    import { Toast } from '$lib/index'
    import { popover } from "$lib/store";
    import { formatDate } from "$lib";

    export let data = {}
    let item

    async function incrementEpisodes(id) {
        let url = `https://api.myanimelist.net/v2/anime/${id}/my_list_status`

        let formData = {
            num_watched_episodes: data.list_status.num_episodes_watched + 1
        }

        if(formData.num_watched_episodes == data.node.num_episodes) {
            formData.status = "completed"
        }

        const response = await PatchRequest(url, JSON.stringify(formData))

        if(!response.success) {
            Toast.fire({
                icon: 'error',
                title: 'Something went wrong!',
            })
            return
        }

        const responseData = JSON.parse(response.data)

        if(responseData.status == "completed") {
            Toast.fire({
                icon: 'success',
                title: 'Anime completed!',
            }).then(() => {
                item.remove()
            })
        }

        data.list_status.num_episodes_watched = responseData.num_episodes_watched
    }

    /**
     * @param event {MouseEvent}
     */
    function handleEdit(event) {
        $popover = {
            ...$popover,
            title: data.node.title,
            episodesTotal: data.node.num_episodes,
            episodesWatched: data.list_status.num_episodes_watched,
            status: data.list_status.status,
            score: data.list_status.score,
            rewatching: data.list_status.is_rewatching.toString(),
            timesRewatched: data.list_status.num_times_rewatched,
            animeId: data.node.id,
        }

        $popover.open()
    }
</script>

<style lang="less">

    .item {
        --_border-raduis: 5px;
        width: calc(250px * var(--your-list-card-size-multiplier, 1));
        height: calc(350px * var(--your-list-card-size-multiplier, 1));
        display: flex;
        justify-content: center;
        align-items: center;
        position: relative;
        border-radius: var(--_border-raduis);
        cursor: pointer;

        transition: transform .1s ease-in-out;

        img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            border-radius: var(--_border-raduis);

            transition: filter .1s ease-in-out;
        }

        .overlay {
            width: 100%;
            height: 100%;
            position: absolute;
            top: 0;
            left: 0;

            display: flex;
            flex-direction: column;

            .top {
                flex-grow: 1;
                padding: .5rem;
                display: flex;
                flex-direction: column;
                gap: .5rem;
                border-top-right-radius: var(--_border-raduis);
                border-top-left-radius: var(--_border-raduis);
                color: var(--text-white);
                background-color: rgba(0, 0, 0, 0.4);
                
                opacity: 0;
                user-select: none;
                pointer-events: none;

                transition: opacity .1s ease-in-out, background .1s ease-in-out, color .1s ease-in-out;

                .row {
                    width: 100%;
                    transform: scale(0.9) translate(0, 1rem);
                    transition: transform .1s ease-in-out;


                    p {
                        display: flex;
                        align-items: center;
                        margin: 0;
                        padding: 0;
                        font-size: 1rem;
                        line-height: 1rem;

                        span {
                            padding-right: .25rem;
                        }
                    }
                }

                .row.one {
                    display: flex;
                    justify-content: space-between;

                    button {
                        all: unset;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        cursor: pointer;

                        transition: all .1s ease-in-out;
                    }

                    button:hover {
                        color: var(--text-white);
                        font-weight: 600;
                        transform: scale(1.1);
                    }
                }
            }

            .bottom {
                width: 100%;
                height: 5rem;
                padding: .5rem;
                background: rgba(0, 0, 0, 0.6);
                color: var(--text-light);
                border-bottom-left-radius: var(--_border-raduis);
                border-bottom-right-radius: var(--_border-raduis);
                display: flex;
                flex-direction: column;

                transition: background .1s ease-in-out, color .1s ease-in-out;

                .title {
                    font-size: 1rem;
                    line-height: 1rem;
                    height: 2rem;
                    line-clamp: 2;
                    margin: 0;
                    padding: 0;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: wrap;
                }

                .split {
                    display: flex;
                    justify-content: space-between;
                    align-items: end;
                    flex-grow: 1;
                    
                    p {
                        margin: 0;
                        padding: 0;
                    }

                    .right {
                        .eps {
                            .plus {
                                position: relative;

                                .hover {
                                    display: flex;
                                    opacity: 0;
                                    user-select: none;
                                    pointer-events: none;
                                }

                                .nohover {
                                    opacity: 1;
                                    user-select: none;
                                    pointer-events: all;
                                }

                                .hover, .nohover {
                                    width: 100px;
                                    height: 21.6px;

                                    display: flex;
                                    justify-content: end;

                                    position: absolute;
                                    right: 0;
                                    bottom: 0;
                                }

                                button {
                                    all: unset;
                                    display: flex;
                                    justify-content: center;
                                    align-items: center;
                                    padding-left: .4rem;
                                    padding-right: .2rem;
                                    
                                    span {
                                        font-size: 1.2rem;
                                        line-height: 1.2rem;
                                        height: 1rem;
                                        margin: 0;
                                        padding: 0;
                                        color: var(--text-light);
                                        cursor: pointer;
                                    }
                                }

                                button:hover {
                                    span {
                                        color: var(--text-white);
                                        font-weight: 600;
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    .item:hover {
        transform: scale(1.05);

        img {
            filter: blur(2px);
        }
            
        .overlay {

            .top {
                opacity: 1;
                user-select: none;
                pointer-events: all;

                .row {
                    transform: scale(1) translate(0, 0);
                }
            }

            .bottom {
                background: rgba(0, 0, 0, 0.8);
                color: var(--text-white);

                .split .eps .plus {
                    .hover {
                        opacity: 1;
                        user-select: none;
                        pointer-events: all;
                    }

                    .nohover {
                        opacity: 0;
                        user-select: none;
                        pointer-events: none;
                    }
                }
            }

        }
    }

</style>