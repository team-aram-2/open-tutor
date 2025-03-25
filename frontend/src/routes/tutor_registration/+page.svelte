<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	$: submittingForm = false;
	$: submissionError = '';
	$: agreementChecked = false;
	$: bankRoutingNumber = '';
	$: bankAccountNumber = '';
	$: dob = '';
	$: phoneNumber = '';
	$: addressLine1 = '';
	$: addressCity = '';
	$: addressState = '';
	$: addressZipCode = '';
	$: ssn = '';

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		submissionError = params.get('err') ?? '';

		// TODO: NEED STATE CHECK HERE FOR LOGGED_IN...
		// This should only be available to users who have logged in and are not already registered as a tutor.
		// Redirect to tutor dashboard if this page is reached by a valid tutor.
		// Redir to /login if the user isn't logged in.
	});
</script>

<div class="w-full flex flex-col text-white overflow-y-scroll mt-20 pb-20">
	<div class="flex flex-col gap-1 w-120 mx-auto">
		<h1 class="text-center text-5xl py-4">Tutor Registration</h1>
		<h2 class="text-2xl py-2">
			<b>Code of Conduct:</b>
		</h2>
		<div class="flex flex-col gap-y-3">
			<b>As a tutor, you agree to:</b>
			<p>
				1. <b>Conduct yourself professionally: </b> A good tutor is understanding, respectful, and acts
				ethically.
			</p>
			<p>
				2. <b>Communicate well.</b> Students rely on you to provide accurate information, and communicate
				with patience and fairness.
			</p>
			<p>
				3. <b>Be kind.</b> Harassment, discrimination, or misconduct of any kind is strictly prohibited.
				Failure to adhere to these standards may result in suspension or termination of your tutoring
				privileges.
			</p>
			<p>
				4. <b>Lead the community.</b> Tutors should never condone or promote academic dishonesty of any
				kind. Tutors will always report instances where a student has engaged in academic dishonesty
				on OpenTutor.
			</p>
			<p>
				By agreeing to the above, you can be expected to be held accountable on these terms-- as
				well as the terms of the <a
					class="underline text-purple-300 hover:text-orange-300"
					href="#top">general OpenTutor user agreement</a
				>.
			</p>
		</div>
		<form class="flex flex-col gap-3" method="POST" action="{PUBLIC_API_HOST}/tutor">
			<div
				class="flex flex-col justify-left items-center gap-4 p-4 my-2 border border-black-200 rounded-sm dark:border-gray-700"
			>
				<div>
					<input
						id="agreement-checkbox"
						class="bg-gray-800 border border-gray-700 rounded-md shadow-inner
					checked:bg-blue-600 checked:border-blue-500 transition-all duration-200
					hover:ring-2 hover:ring-blue-400 focus:ring-2 focus:ring-blue-400 cursor-pointer"
						type="checkbox"
						bind:checked={agreementChecked}
					/>
					<label for="agreement-checkbox" class="w-2/3">I agree to the above.</label>
				</div>

				<!-- Personal info -->
				<div>
					<label for="dob" class="mr-4">Date of Birth</label>
					<input
						id="dob"
						name="dob"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type="date"
						placeholder=""
						bind:value={dob}
					/>
				</div>
				<div>
					<label for="phone" class="mr-4">Phone number</label>
					<input
						id="phone"
						name="phone"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type="tel"
						placeholder=""
						bind:value={phoneNumber}
					/>
				</div>

				<hr class="border b-2 w-full border-gray-600" />

				<!-- Address -->
				<div>
					<label for="address_line1" class="mr-4">Street address</label>
					<input
						id="address_line1"
						name="address_line1"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder=""
						bind:value={addressLine1}
					/>
				</div>
				<div>
					<label for="address_city" class="mr-4">City</label>
					<input
						id="address_city"
						name="address_city"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder=""
						bind:value={addressCity}
					/>
				</div>
				<div>
					<label for="address_state" class="mr-4">State</label>
					<input
						id="address_state"
						name="address_state"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder=""
						bind:value={addressState}
					/>
				</div>
				<div>
					<label for="address_postalcode" class="mr-4">Zip code</label>
					<input
						id="address_postalcode"
						name="address_postalcode"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder=""
						bind:value={addressZipCode}
					/>
				</div>

				<hr class="border b-2 w-full border-gray-600" />

				<!-- Banking info -->
				<div>
					<label for="bank_routing_number" class="mr-4">Routing number</label>
					<input
						id="bank_routing_number"
						name="bank_routing_number"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder="Bank routing number"
						bind:value={bankRoutingNumber}
					/>
				</div>
				<div>
					<label for="bank_account_number" class="mr-4">Account number</label>
					<input
						id="bank_account_number"
						name="bank_account_number"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						placeholder="Bank account number"
						bind:value={bankAccountNumber}
					/>
				</div>
				<div>
					<label for="ssn" class="mr-4">Full tax ID (SSN)</label>
					<input
						id="ssn"
						name="ssn"
						class="bg-gray-900 px-4 py-2 rounded-md"
						type=""
						maxlength="9"
						placeholder=""
						bind:value={ssn}
					/>
				</div>

				<input
					class="rounded-md bg-blue-600 w-full py-2 enabled:cursor-pointer disabled:opacity-50"
					type="submit"
					value="Register"
					disabled={submittingForm ||
						!agreementChecked ||
						!dob ||
						!phoneNumber ||
						!addressLine1 ||
						!addressCity ||
						!addressState ||
						!addressZipCode ||
						!ssn ||
						!bankAccountNumber ||
						!bankRoutingNumber}
				/>
			</div>
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>

<style>
	input[id='agreement-checkbox'] {
		transform: scale(1.5);
		margin: 10px;
	}
</style>
