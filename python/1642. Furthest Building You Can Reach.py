import heapq
def furthestBuilding(heights, bricks, ladders):
    n = len(heights)
    heap = []

    for i in range(n - 1):
        h = heights[i + 1] - heights[i]

        if h <= 0:
            continue

        heapq.heappush(heap, h)
        print(heap)

        if len(heap) > ladders:
            min_h = heapq.heappop(heap)
            bricks -= min_h

        if bricks < 0:
            return i

    return n - 1



heights = [4,6,7,6,9,14,12]
bricks = 1
ladders = 1


res = furthestBuilding(heights, bricks, ladders)

print(res)