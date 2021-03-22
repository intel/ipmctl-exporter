set(GOPATH "${CMAKE_CURRENT_BINARY_DIR}/go")
file(MAKE_DIRECTORY ${GOPATH})

function(add_go_executable NAME)
    file(GLOB GO_SOURCE RELATIVE "${CMAKE_CURRENT_SOURCE_DIR}" "*.go")
    add_custom_command(OUTPUT ${OUTPUTDIR}/.timestamp
        COMMAND ${CMAKE_Go_COMPILER} build
        -ldflags "-X main.Version=${IPMCTL_EXPORTER_VERSION_STRING}"
        -mod=mod
        -o "${CMAKE_CURRENT_BINARY_DIR}/${NAME}.exe"
        ${CMAKE_GO_FLAGS} ${GO_SOURCE}
        WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})

    add_custom_target(${NAME} ALL DEPENDS ${OUTPUTDIR}/.timestamp ${ARGN})
    install(PROGRAMS ${CMAKE_CURRENT_BINARY_DIR}/${NAME}.exe DESTINATION bin COMPONENT ipmctl_exporter)
endfunction(add_go_executable)
