import { tob62 } from "@/lib/base62";
import { db } from "@/lib/db";
import { snowflake } from "@/lib/snowflake";
import { redirect } from "next/navigation";

export default async function Page({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  let { url } = await searchParams;

  if (!url) {
    return redirect("/");
  }

  // searchParams can be composed into an array if the same key is defined multiple times in the request URL.
  // Should be relatively safe to index the array here, as it takes at least 2 elements for `url` to come through as an array.
  if (Array.isArray(url)) {
    url = url[0];
  }

  if (!isValidHttpUrl(url)) {
    throw "Invalid URL";
  }

  const results = await getShortUrl(url);

  return <p>{results}</p>;
}

function isValidHttpUrl(string: string): boolean {
  let url: URL;

  try {
    url = new URL(string);
  } catch {
    return false;
  }

  return url.protocol === "http:" || url.protocol === "https:";
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
