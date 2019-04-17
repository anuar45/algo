def bfs(graph, root):
    visited, queue = [], [root]
    while queue:
        vertex = queue.pop(0)
            for w in graph[vertex]:
                if w not in visited:
                    visited.append(w)
                    queue.append(w)

graph = {0: [1,2], 1: [2], 2: []}
bfs(graph, 0)
