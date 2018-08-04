package helloworld

const HELLOWORLD = "hello wworld"

func CreateMessage(prefix string) string {
    if prefix == "" {
        return "NO PREFIX: " + HELLOWORLD
    }

    return prefix + HELLOWORLD
}
