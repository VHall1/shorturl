import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Form from "next/form";

export default function Home() {
  return (
    <Form action="/shortener">
      <div className="grid w-full max-w-sm items-center gap-1.5">
        <Label htmlFor="email">URL</Label>
        <Input name="query" placeholder="https://example.com" />
      </div>

      <Button type="submit" className="mt-4">Shorten URL</Button>
    </Form>
  );
}
