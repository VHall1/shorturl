import { tob62 } from "@/lib/base62";
import { db } from "@/lib/db";
import { snowflake } from "@/lib/snowflake";

export default async function Page({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  const results = await getShortUrl((await searchParams).query);

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
