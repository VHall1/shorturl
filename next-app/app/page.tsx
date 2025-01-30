import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Form from "next/form";

export default function Page() {
  return (
    <Form action="/shorten">
      <div className="flex flex-col lg:flex-row gap-2">
        <Input
          name="url"
          type="url"
          placeholder="Paste your long URL here"
          className="flex-1"
          required
        />
        <Button type="submit">Shorten URL</Button>
      </div>
    </Form>
  );
}
