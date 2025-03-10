import type { PageLoad } from './$types';
import { PUBLIC_API_HOST } from '$env/static/public';

export const load: PageLoad = async ({ fetch, params }) => {
	// const { id } = params;
  console.log("params: " + String(params));

  const url = PUBLIC_API_HOST + "/meetings";
  console.log("url: " + url);
  // Fetch all meetings for user
	await fetch(
    url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }
  )
  // Check for error retrieving appointments
  .then(response => {
    if (!response.ok){
      throw new Error("HTTP error. Status code: " + String(response.status));
    }
    return response.json();
  })
  // Sort through response to extract unique tutors
  .then(data => {
    console.log(data)
  })
  .catch(error => {
    console.error("Something went wrong. ", error);
  });

};
