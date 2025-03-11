import Form from "next/form";

export default function Home() {
  return (
    <Form action="/s">
      <input name="query" />
      <button type="submit">Submit</button>
    </Form>
  );
}
