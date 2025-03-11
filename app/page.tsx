import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Form from "next/form";

export default function Home() {
  return (
    <Form action="/shorten">
      <div className="flex flex-col lg:flex-row gap-2">
        <Input
          name="url"
          type="url"
          placeholder="https://example.com"
          className="flex-1"
          required
        />
        <Button type="submit">Shorten URL</Button>
      </div>
    </Form>
  );
}
