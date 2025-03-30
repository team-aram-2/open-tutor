<script lang="ts">
	export let skill: string;
	export let selectedSkills: string[];
	export let availableSkills: string[];
	import { toggleSkill } from './tutor-search';
</script>

<button
	id={skill}
	class="skill"
	class:selected={selectedSkills.includes(skill)}
	on:click={() => {
		toggleSkill(skill);
	}}>{skill}</button
>
