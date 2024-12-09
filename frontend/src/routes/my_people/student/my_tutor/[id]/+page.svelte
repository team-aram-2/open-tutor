<script lang="ts">
    import TutorPage from '$lib/components/cards/tutor-page.svelte'
    import tutorsData from '$lib/mock/my_tutors_mock.json'
    export let data: { user: { firstName: string; lastName: string; userId: string; skills: string[]; overallRating: number;} };
</script>

<div>
    <div class="cardcontainer">
        <TutorPage
        name="{data.user.firstName} {data.user.lastName}"
        rating={data.user.overallRating}
        skills={data.user.skills}
        userId="{data.user.userId}"
        ></TutorPage>
    </div>
</div>

<style>
    .cardcontainer{
        display:flex;
        flex-direction: row;
        flex-wrap:wrap;
        height: 100%;

    }

</style>