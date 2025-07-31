CREATE TABLE robots (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    model VARCHAR(100) NOT NULL,
    version VARCHAR(50),
    color VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
INSERT INTO robots (name, description, model, version, color)
VALUES 
('Robot Alpha', 'Robot dùng trong nghiên cứu AI', 'RX-01', 'v1.0', 'Đỏ'),
('Robot Beta', 'Robot vệ sinh tự động trong nhà', 'RX-02', 'v2.3', 'Xanh dương'),
('Robot Gamma', 'Robot vận chuyển hàng trong kho', 'RX-03', 'v1.5', 'Vàng'),
('Robot Delta', 'Robot hỗ trợ y tế từ xa', 'RX-04', 'v3.0', 'Trắng'),
('Robot Epsilon', 'Robot dùng trong giáo dục lập trình', 'RX-05', 'v1.1', 'Đen');
