import { getRedirectUrl } from "@/lib/url-shortener";
import { redirect, RedirectType } from "next/navigation";

export async function GET(
  _: Request,
  { params }: { params: Promise<{ url: string }> }
) {
  const url = await getRedirectUrl((await params).url);
  return redirect(url, RedirectType.push);
}
