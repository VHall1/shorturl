import { db } from "@/lib/db";
import { redirect, RedirectType } from "next/navigation";

export async function GET(
  _: Request,
  { params }: { params: Promise<{ id: string }> }
) {
  const url = await db.url.findFirst({
    where: { shortUrl: (await params).id },
  });

  return redirect(url?.longUrl ?? "/", RedirectType.replace);
}
