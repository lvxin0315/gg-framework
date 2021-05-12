package examples

import (
	"fmt"
	"testing"
)

//	 ####   ####        #####   ####   ####  #    #
//	#    # #    #       #    # #    # #    # ##  ##
//	#      #      ##### #    # #    # #    # # ## #
//	#  ### #  ###       #####  #    # #    # #    #
//	#    # #    #       #   #  #    # #    # #    #
//	 ####   ####        #    #  ####   ####  #    #

func TestComment(t *testing.T) {
	c := ` ####   ####        #####   ####   ####  #    # 
#    # #    #       #    # #    # #    # ##  ## 
#      #      ##### #    # #    # #    # # ## # 
#  ### #  ###       #####  #    # #    # #    # 
#    # #    #       #   #  #    # #    # #    # 
 ####   ####        #    #  ####   ####  #    # 
                                                `
	fmt.Print(c)
}