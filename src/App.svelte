<script lang="ts">
import { ServiceClient } from "./proto/service.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

let content = 'Click "Request" to fetch the backend.';

const client = new ServiceClient(
  new GrpcWebFetchTransport({
    baseUrl: `http://${window.location.hostname}:7000`,
  })
);
</script>

<main class="p-4">
  <p class="my-2">{content}</p>
  <button
    class="border border-neutral-800 bg-neutral-200 rounded-md py-1 px-2"
    on:click={() => {
      client.echo({}).then(({ response }) => {
        content = response.content;
      });
    }}
  >
    Request
  </button>
</main>
