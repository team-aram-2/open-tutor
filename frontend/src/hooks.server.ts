import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({event, resolve}) => {
  console.log(event.url.pathname);
  const protectedRoutes: string[] = ['/messages/student'];
  if(protectedRoutes.includes(event.url.pathname)) {
    const userCookie: string | undefined = event.cookies.get('session_token');
    console.log(userCookie);
    if (!userCookie) {
      throw redirect(303, `/login?redirectTo=${event.url.pathname}`)
    }
  }
  return resolve(event)
}