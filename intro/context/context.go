/*
time.Sleep blocks the gather data function and context,
so even after gather data ctx get timeout error after 1 seconds,
it does not return error before 7 seconds sleep,
after 7 seconds sleep gather data context get timeout error and it \
	and it propragte to child context localContext
	and when it enter getFilesLocal ctx.Err is not null and return error

if global timeout is 4 seconds and GatherData sleep is 2 second then it blocks until the
getFilesLocal 15 minute sleep complete

** it is better to use loop to check ctx timeout
*/
package main
import (
	"fmt"
	"context"
	"time"
)
type Args struct{
	local []int
}
func getFilesLocal(ctx context.Context, args []int) ([]int, error){
	fmt.Println("starting local files")
	if ctx.Err() != nil{
		return nil, ctx.Err()
	}
	time.Sleep(15 * time.Second)
	fmt.Println("finish local files sleep")
	args = append(args, 4)
	return args, nil
}
func GatherData(ctx context.Context, args Args) ([]int, error){
	fmt.Println("starting gather context")
	if ctx.Err() != nil{
		return nil, ctx.Err()
	}
	// select {
	// case <- ctx.Done():
	// 	return nil, ctx.Err()
	// }
	args.local = append(args.local, 7)
	time.Sleep(7*time.Second)
	fmt.Println("gather data sleep finished")
	localCtx, localCancel := context.WithTimeout(ctx, 2*time.Second)
	defer localCancel()
	local, err := getFilesLocal(localCtx, args.local)
	if err != nil {
		fmt.Println("Found error in local function: ", err)
		return nil, err
	}
	fmt.Println("getting updated file: ", local)
	
	return args.local, nil
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer func(){
		fmt.Println("cancelling context")
		cancel()
	}()
	args := Args{local: []int{1, 2, 3}}
	data, err := GatherData(ctx, args)
	if err != nil{
		fmt.Println("Gather data ran into error: ", err)
	}
	fmt.Println(data)
}