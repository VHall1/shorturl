"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { useRef, useState } from "react";

export function ShortUrlBox({ shortUrl }: { shortUrl: string }) {
  const $input = useRef<HTMLInputElement>(null);
  const [isCopyDone, setIsCopyDone] = useState(false);

  const handleCopy = async () => {
    $input.current?.select();
    $input.current?.setSelectionRange(0, 99999); // for mobile

    // TODO: add state to change button style once copy action is done?
    await navigator.clipboard.writeText($input.current?.value ?? "");
    setIsCopyDone(true);
  };

  let copyBtnText = "Copy URL";
  if (isCopyDone) {
    copyBtnText = "Link copied";
  }

  return (
    <div className="flex flex-col lg:flex-row gap-2">
      <Input defaultValue={shortUrl} className="flex-1" ref={$input} />
      {/* TODO: hide button if js is disabled? */}
      <Button
        onClick={handleCopy}
        className={cn({ "bg-green-500 hover:bg-green-500/90": isCopyDone })}
      >
        {copyBtnText}
      </Button>
    </div>
  );
}
