import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Form from "next/form";

export default function Home() {
  return (
    <Card className="max-w-md mx-auto">
      <CardContent>
        <Form action="/shorten">
          <div className="grid w-full max-w-sm items-center gap-1.5">
            <Label htmlFor="url">URL</Label>
            <Input id="url" name="url" placeholder="https://example.com" />
          </div>

          <Button type="submit" className="mt-4 w-full lg:w-auto">
            Shorten URL
          </Button>
        </Form>
      </CardContent>
    </Card>
  );
}
