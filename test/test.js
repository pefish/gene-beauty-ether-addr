
function main (args) {
    var address = args[0]
    return regex.match("^0x(.)\\1{5,42}.*", address)
    || regex.match("^0xpefish.*", address)
    || regex.match("^0x123456.*", address)
    || regex.match("^0x654321.*", address)
    || regex.match("^0x012345.*", address)
    || regex.match("^0x543210.*", address)
    // || regex.match("^0xa.*", address)
}
