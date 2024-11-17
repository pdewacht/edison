#!/usr/bin/expect

proc simh args {
    # Enter a command on the SIMH prompt
    send "[join $args]\n"
    expect "sim>"
}

proc simh_break {} {
    # Interrupt simulation and return to SIMH
    send "\005"
    expect "sim>"
}

proc edison args {
    # Enter a command on the Edison prompt
    foreach line $args {
        send "$line\r\n"
    }
    expect "Command ="
}

proc configure_machine {} {
    # Configure SIMH for the Edison system
    simh set cpu 11/23
    simh set cpu 64k
    simh set rx disabled
    simh set ry enabled
    simh set ry0 single
    simh set ry1 single

    # SIMH doesn't support booting single-density floppies
    # with the ry controller. Work around that by manually
    # entering the bootloader.
    set address 0
    foreach word [bootloader] {
        simh deposit [format "%03ho %06ho" $address $word]
        incr address 2
    }
}

proc bootloader {} {
    set f [open bin-edison/kernel]
    fconfigure $f -translation binary
    binary scan [read $f 128] s64 words
    close $f
    return $words
}

proc boot {} {
    send "go 0\n"
    expect "The Edison system"

    expect "b if both disks are blank"
    send "s"
    expect "Command ="
}

proc swap_disks {fmt attach_commands} {
    send "insert\r\n"
    expect "b if both disks are blank"
    simh_break
    uplevel $attach_commands
    send "cont\n"
    sleep 0.1; # why is this needed?
    send $fmt
    if {$fmt ne "s"} {
        send "\r\n"
    }
    expect "Command ="
}

spawn pdp11

#-re {line[ \a]*[0-9][^\n]*}  abort
expect_before {
    default                      abort
    -re {line[ \a]*[0-9][^\n]*}  abort
    "compilation errors"         ftbfs
    "Files are different"        abort
}
proc abort {} {
    puts "\nLooks like something went wrong :("
    exit 1
}
proc ftbfs {} {
    expect "Command ="
    edison type notes
    abort
}

expect "sim>"
configure_machine
file copy -force disks/bootstrap.dsk disks/tmp.dsk
simh attach ry0 disks/tmp.dsk
simh attach ry1 disks/sources1.dsk
boot

puts {
    #
    # Build a few tools
    #
}

edison compile comparetext compare 0 no
edison compile typetext type 0 no
edison compile dumptext dump 0 no

puts {
    #
    # Re-build the Edison compiler
    #
}

edison compile compiletext newcompile 0 no
edison compile edison1text newedison1 0 no
edison compile edison2text newedison2 0 no
swap_disks s {
    simh attach ry1 disks/sources2.dsk
}
edison compile edison3text newedison3 0 no
edison compile edison4text newedison4 0 no

puts {
    #
    # Install the newly-built compiler
    #
}

edison delete 0 compile
edison delete 0 edison1
edison delete 0 edison2
edison delete 0 edison3
edison delete 0 edison4
edison rename 0 newcompile compile
edison rename 0 newedison1 edison1
edison rename 0 newedison2 edison2
edison rename 0 newedison3 edison3
edison rename 0 newedison4 edison4

puts {
    #
    # Now check if it works:
    #   build the compiler again
    #
}

swap_disks s {
    simh attach ry1 disks/sources1.dsk
}
edison compile compiletext newcompile 0 no
edison compile edison1text newedison1 0 no
edison compile edison2text newedison2 0 no
swap_disks s {
    simh attach ry1 disks/sources2.dsk
}
edison compile edison3text newedison3 0 no
edison compile edison4text newedison4 0 no

puts {
    #
    # Check that the binaries match
    #
}

edison compare compile newcompile
edison compare edison1 newedison1
edison compare edison2 newedison2
edison compare edison3 newedison3
edison compare edison4 newedison4

puts {
    #
    # Re-build the system
    #
}

swap_disks s {
    simh attach ry1 disks/sources1.dsk
}
edison compile systemtext system 0 no
edison compile typetext type 0 no
edison compile dumptext dump 0 no
edison compile comparetext compare 0 no

puts {
    #
    # Re-build the kernel
    #
}

swap_disks s {
    simh attach ry1 disks/sources2.dsk
}
edison compile alvatext alva 0 no

swap_disks s {
    simh attach ry1 disks/sources3.dsk
}
edison alva kerneltext kernel 0

puts {
    #
    # Make a new system disk
    #
}

swap_disks 0 {
    simh attach ry1 -n disks/system.dsk
}
edison newkernel 1 kernel
edison newsystem 1 system
edison copy kernel kernel 1
edison copy system system 1
edison copy compile compile 1
edison copy edison1 edison1 1
edison copy edison2 edison2 1
edison copy edison3 edison3 1
edison copy edison4 edison4 1
edison copy alva alva 1
edison copy type type 1
edison copy dump dump 1
edison copy compare compare 1

edison list 1

puts {
    #
    # All done
    #
}

simh_break
send "bye\n"
expect "Goodbye"
wait
file delete disks/tmp.dsk
