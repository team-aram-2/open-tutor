<script>
    export let profilePicture = 'git/images/cards/pfp.png';
    export let name = 'John Doe';
    export let rating = 2;
    export let starImage = "/images/cards/star.png";
    export let emptyStarImage = "/images/cards/empty_star.png";
</script>

<div class="tutor-card flex h-screen">
    <h1 class="tutor-header">
        <img class="pfp" src={profilePicture} alt=""/>
        <h2 class="tutor-name">{name}
            <div class="rating">
                {#each Array(5) as _, index}
                    <img class="star" 
                        src={index < rating ? starImage : emptyStarImage}
                        alt="">
                {/each}  
            </div>
        </h2>
    </h1>
</div>

<style>
    
    .tutor-card {

        padding-right: 20px;
        padding-bottom: 100px;
        margin-bottom: 25px;

        border-top-right-radius: 25px;
        border-bottom-right-radius: 25px;
        border-top-left-radius: 25px;
        border-bottom-left-radius: 25px;

        min-width: 500px;
        min-height: 375px;
        height: auto;
        width: auto;
        max-height: 1000px;


        display: flex;
        flex-direction: column;

        background: #FDF8ED;
        border: 5px solid #EBAD33;
    }

    .tutor-header {
        display: flex;
        flex-direction: row;
    }
    .tutor-name {
        display: flex;
        flex-direction: column;
        
        position: relative;
        top: 0;
        left: 0;

        padding-left: 25px;
        padding-right: 25px;
        margin-top: 10px;

        max-height: 50px;
        margin-bottom: 0;

        font-size: 50px;
        font-weight: bold;
        color: #352505;
        /* border: 2px solid blue */
    }
    .pfp {
        width: 120px;
        height: 120px;

        margin-left: 20px;


        border-radius: 15px;
        border-top-right-radius: 25px;
        border-bottom-right-radius: 25px;
        border-top-left-radius: 25px;
        border-bottom-left-radius: 25px;

    }

    .rating{
        display: flex;
        flex-direction: row;

    }

    .star{
        width: 40px;
        height: 40px;
        margin-left: 5px;
    }

    .no-decoration{
        text-decoration: none;
        color: inherit;
    }
</style>