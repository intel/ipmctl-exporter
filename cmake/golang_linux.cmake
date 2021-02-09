set(GOPATH "${CMAKE_CURRENT_BINARY_DIR}/go")
file(MAKE_DIRECTORY ${GOPATH})

function(add_go_executable NAME)
    file(GLOB GO_SOURCE RELATIVE "${CMAKE_CURRENT_SOURCE_DIR}" "*.go")
    execute_process(
        COMMAND sh -c "go version | grep -o [0-9].[0-9]*.[0-9]* "
        OUTPUT_VARIABLE currentver
    )
    set(ENV{CURRENTGOVER} "${currentver}")
    # -- REQUIREDGOVER - version number since -mod=mod is supported
    set(ENV{REQUIREDGOVER} "1.14.0")
    message("GO VERSION: " $ENV{CURRENTGOVER})
    execute_process(
        COMMAND sh -c "if [ $(printf '%s\n' $REQUIREDGOVER $CURRENTGOVER | sort -V | head -n1) = $REQUIREDGOVER ]; then echo 1; else echo 0; fi"
        OUTPUT_VARIABLE isGOVersionNewerThanRequired
    )
    # -- build -mod=mod is supported only since go version 1.14.x, older versions have no option like that - hence build would fail
    if(${isGOVersionNewerThanRequired} EQUAL 1)    
        add_custom_command(OUTPUT ${OUTPUTDIR}/.timestamp
            COMMAND env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build
            -ldflags "-X main.Version=${IPMCTL_EXPORTER_VERSION_STRING}"
            -mod=mod
            -o "${CMAKE_CURRENT_BINARY_DIR}"
            ${CMAKE_GO_FLAGS} ${GO_SOURCE}
            WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
    else()
        add_custom_command(OUTPUT ${OUTPUTDIR}/.timestamp
            COMMAND env GOPATH=${GOPATH} ${CMAKE_Go_COMPILER} build
            -ldflags "-X main.Version=${IPMCTL_EXPORTER_VERSION_STRING}"
            -o "${CMAKE_CURRENT_BINARY_DIR}"
            ${CMAKE_GO_FLAGS} ${GO_SOURCE}
            WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
    endif()

    add_custom_target(${NAME} ALL DEPENDS ${OUTPUTDIR}/.timestamp ${ARGN})
    install(PROGRAMS ${CMAKE_CURRENT_BINARY_DIR}/${NAME} DESTINATION bin)
endfunction(add_go_executable)