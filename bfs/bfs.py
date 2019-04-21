def bfs(graph, root):
    visited, queue = [], [root]
    while queue:
        vertex = queue.pop(0)
        for w in graph[vertex]:
            if w not in visited:
                print(w)
                visited.append(w)
                queue.append(w)

graph = {0: [1,2], 1: [3,4], 2: [5,6], 3: [7,8], 4: [9,10]}
bfs(graph, 0)
