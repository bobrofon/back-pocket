cmake_minimum_required (VERSION 3.5)

project (back-pocket NONE)

enable_testing()
add_test(
	NAME back-pocket-test
	COMMAND ${CMAKE_CURRENT_SOURCE_DIR}/scripts/test.sh)

add_custom_command(
	OUTPUT ${PROJECT_NAME}
	#COMMAND go build ${CMAKE_CURRENT_SOURCE_DIR}/cmd/${PROJECT_NAME})
	COMMAND go build github.com/bobrofon/${PROJECT_NAME}/cmd/${PROJECT_NAME})

add_custom_target(build ALL DEPENDS ${PROJECT_NAME})
