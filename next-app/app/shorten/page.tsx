import { Button } from "@/components/ui/button";
import { shortenUrl } from "@/lib/url-shortener";
import Link from "next/link";
import { ShortUrlBox } from "./short-url-box";

export default async function Shorten({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | undefined }>;
}) {
  const longUrl = (await searchParams).url ?? "";
  const shortUrl = await shortenUrl(longUrl);

  return (
    <div>
      <ShortUrlBox shortUrl={shortUrl} />
      <Button variant="outline" className="mt-2 w-full lg:w-auto" asChild>
        <Link href="/">Shorten another URL</Link>
      </Button>

      <div className="mt-4">
        Long URL:
        <Link
          className="text-primary underline-offset-4 hover:underline ml-1"
          href={longUrl}
        >
          {longUrl}
        </Link>
      </div>
    </div>
  );
}
