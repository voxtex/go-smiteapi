# Go Smite API Client

## Usage

    client := smiteapi.NewClient("my-auth-key", "my-dev-id")

    // Invoke API methods
    client.GetDataUsed()
    client.GetItems()

    // Useful Constants
    smiteapi.Conquest5v5
    smiteapi.Bronze

    // Change language
    client.SetLanguage(smiteapi.French)

- Sessions will automatically be created when needed.
- Client should be thread safe.
- The JSON response names are normalized - Active_Item_1 -> ActiveItem1.
- Does not currently support all API methods.
- Constants are exposed for easy use
