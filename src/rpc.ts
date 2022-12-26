import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";

const managementState: { [key: symbol]: boolean } = {}

export function manageStream<T extends object>(
    generate: () => ServerStreamingCall<any, T>,
    onMessage: (message: T) => void,
    key: symbol | undefined = undefined
): () => void {
    if (!key) {
        key = Symbol("managed-stream")
    }
    managementState[key] = true

    const call = generate()
    call.responses.onMessage(onMessage)
    call.responses.onError(e => {
        if (!e.message.includes("NetworkError")) {
            return
        }
        setTimeout(() => {
            if (!managementState[key!]) {
                return
            }
            manageStream(generate, onMessage, key);
        }, 3000)
    })

    return () => {
        delete managementState[key!]
    }
}
