<div class="item">

    <div class="left">
        <img src={data.node.main_picture.medium} alt={data.node.title} draggable="false" loading="lazy">
    </div>

    <div class="right">
        <h2>{data.node.title}</h2>
    </div>

</div>

<script>
    export let data = {}
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

            h2 {
                margin: 0;
                padding: 0.5rem;
                font-size: 1.5rem;
                font-weight: 400;
                overflow-x: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
        }
    }

</style>