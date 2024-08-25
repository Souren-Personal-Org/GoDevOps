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
	if ctx.Err() != nil{
		return nil, ctx.Err()
	}
	time.Sleep(10 * time.Second)
	args = append(args, 4)
	return args, nil
}
func GatherData(ctx context.Context, args Args) ([]int, error){
	if ctx.Err() != nil{
		return nil, ctx.Err()
	}
	args.local = append(args.local, 7)
	time.Sleep(6 * time.Second)
	localCtx, localCancel := context.WithTimeout(ctx, 2*time.Second)
	local, err := getFilesLocal(localCtx, args.local)
	if err != nil {
		fmt.Println("Found error in local function: ", err)
		return nil, err
	}
	fmt.Println("getting updated file: ", local)
	localCancel()
	return args.local, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	args := Args{local: []int{1, 2, 3}}
	data, err := GatherData(ctx, args)
	if err != nil{
		fmt.Println("Gather data ran into error: ", err)
	}
	fmt.Println(data)
	cancel()
}