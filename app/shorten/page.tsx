import { tob62 } from "@/lib/base62";
import { db } from "@/lib/db";
import { snowflake } from "@/lib/snowflake";
import { redirect } from "next/navigation";

export default async function Page({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  const { url } = await searchParams;

  if (!url) {
    return redirect("/");
  }

  const results = await getShortUrl(
    // searchParams can be composed into an array if the same key is defined multiple times in the request URL.
    // Should be relatively safe to index the array here, as it takes at least 2 elements for `url` to come through as an array.
    Array.isArray(url) ? url[0] : url
  );

  return <p>{results}</p>;
}

async function getShortUrl(longUrl: string): Promise<string> {
  const url = await db.url.findFirst({ where: { longUrl } });
  if (url) {
    return url.shortUrl;
  }

  const id = snowflake.generate();
  const { shortUrl } = await db.url.create({
    data: { id, longUrl, shortUrl: tob62(id) },
  });

  return shortUrl;
}
