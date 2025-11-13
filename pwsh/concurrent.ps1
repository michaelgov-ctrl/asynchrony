# import
$classes = (Resolve-Path -Path ".\tcp.ps1").ProviderPath
. $classes

# With a listener and a client in the same process, the two
# must overlap, so there is a requirement of concurrency.

$listener = [Listener]::new()
$client = [Client]::new()

# This can not be run serially because it will
# deadlock on Accept forever and the client never runs
$listener.Accept()
$client.Connect() # unreachable



# Since concurrency is required for this workflow to function.
# The code should express that these tasks MUST be progressed concurrently.
# Under the hood ThreadJob could be parallel execution or task switching
Start-ThreadJob -ScriptBlock { 
    . $using:classes
    [Listener]::new().Accept()
}

[Client]::new().Connect()

# This example exhibits asynchrony because it doesn't matter whether the listener
# accepts or the client reaches out first.
# -- There is asynchrony with a dependency on concurrency.