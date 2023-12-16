import random
import networkx as nx
import matplotlib.pyplot as plt

def generate_graph(nodes):
    graph = nx.Graph()
    
    total_connections = nodes * (nodes - 1) // 2
    half_connnections = nodes // 2
    num_connections = random.randint(half_connnections, total_connections)

    connections = set()
    while len(connections) < num_connections:
        u, v = random.sample(range(nodes), 2)
        if u != v and (u, v) not in connections and (v, u) not in connections:
            connections.add((u, v))
    graph_nodes = list(set([i for t in list(connections) for i in t]))
    graph.add_nodes_from(graph_nodes)
    probabilities = [{} for i in range(nodes)]
    for u in graph.nodes():
        u_connections = [v for (x, v) in connections if x == u] + [v for (v, x) in connections if x == u]
        if u_connections:
            probs = [random.uniform(0, 1) for _ in range(len(u_connections))]
            normalization_factor = sum(probs)
            probs = [prob / normalization_factor for prob in probs]
            probabilities[u] = {v: prob for v, prob in zip(u_connections, probs)}

    for u in graph.nodes():
        u_connections = list(probabilities[u].keys())
        for v in u_connections:
            t = random.uniform(0, 1000)
            prob_uv = probabilities[u][v]
            prob_vu = probabilities[v][u]
            graph.add_edge(u, v, p_uv=prob_uv, p_vu=prob_vu, t=t)

    return graph

def save_connections(graph, filename):
    num_nodes = graph.number_of_nodes()
    random_nodes = random.sample(range(num_nodes), 3)
    A,B,C = random_nodes
    while not nx.has_path(graph, A, C) or not nx.has_path(graph, B, C):
        random_nodes = random.sample(range(num_nodes), 3)
        A,B,C = random_nodes
    num_edges = graph.number_of_edges()
    with open(filename, 'w') as file:
        file.write(f"{num_nodes} {num_edges} {C} {A} {B}\n")
        for u in graph.nodes():
            neighbors = list(graph.neighbors(u))
            for v in neighbors:
                data = graph.get_edge_data(u, v)
                prob_uv = data['p_uv']
                prob_vu = data['p_vu']
                t = data['t']
                file.write(f"{u} {v} {t} {prob_uv} {prob_vu}\n")

if __name__ == "__main__":
    nodes_count = random.randint(3, 300)
    generated_graph = generate_graph(nodes_count)
    save_connections(generated_graph, "graph_connections.txt")

    # Draw the graph
    # plt.figure(figsize=(8, 6))
    # pos = nx.spring_layout(generated_graph)
    # nx.draw(generated_graph, pos, with_labels=True, node_size=500, node_color='skyblue', font_weight='bold', font_size=10, arrows=True)
    # edge_labels = {(u, v): f"{data['p_uv']:.2f}, {data['p_vu']:.2f}" for u, v, data in generated_graph.edges(data=True)}
    # nx.draw_networkx_edge_labels(generated_graph, pos, edge_labels=edge_labels, font_color='red')
    # plt.title("Generated Graph")
    # plt.show()
