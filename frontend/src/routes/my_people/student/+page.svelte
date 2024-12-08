<script>
    import Tutor from '$lib/components/cards/tutor-card.svelte';
    import tutorsData from '$lib/mock/my_tutors_mock.json'
</script>

<div>
    <div class="cardcontainer">
        {#each tutorsData.tutors as tutor}
            <Tutor
            name="{tutor.firstName} {tutor.lastName}"
            rating={tutor.overallRating}
            ></Tutor>
        {/each}
    </div>
</div>

<style>
    .cardcontainer{
        display:flex;
        flex-direction: row;
        flex-wrap:wrap;
        gap: 75px;
        flex: 0 0 33%;

    }
</style>