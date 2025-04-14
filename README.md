# Elderwise by Masukin Andre ke Raion

![ElderWise by Masukin Andre ke Raion (1)](https://github.com/user-attachments/assets/73bd0027-17f8-432e-9a31-b3d01ff1192b)

## 🌟 About the Application

Elderwise is an innovative application designed to improve the quality of life for elderly people through technology. Developed by the "Masukin Andre ke Raion" team for the FIND-IT UGM 2025 Hackathon competition, this app connects elderly people with their caregivers in an integrated platform.

## ✨ Main Features

### For Elderly Users (Elder)

- **Elder Mode**: Simple and user-friendly interface for seniors
- **Location Tracking**: Allows caregivers to monitor the elderly's location in real-time
- **SOS/Emergency Button**: Feature to send alerts to caregivers in emergency situations
- **Activity Reminders**: Notifications for medications, activities, and important appointments
- **Fall Detection**: Automatic detection system that alerts caregivers if an elderly person falls

### For Caregivers

- **Monitor Elderly**: View location, activities, and health status of the elderly
- **Agenda Management**: Create and manage activity schedules for the elderly
- **Location History**: Access the elderly's location history to ensure safety
- **Geofence Areas**: Define safe zones and receive alerts if the elderly exits the area
- **Real-time Notifications**: Alerts for emergency conditions or suspicious activities

## 🛠️ Technologies Used

### Frontend

- **Flutter**: Cross-platform application development framework
- **BLoC (Business Logic Component)**: State management pattern for applications

### Backend

- **Golang with Fiber**: Ultra-fast web framework for building APIs
- **PostgreSQL**: Relational database for data storage
- **Swagger**: API documentation and testing

### Cloud & Infrastructure

- **Docker**: Application containerization for consistent development and deployment
- **GitHub Action**: CI/CD pipeline for build and deployment automation
- **DigitalOcean**: Cloud hosting for backend and database
- **Cloudflare**: Web security and performance optimization

### Authentication & Storage

- **Firebase**: Authentication, cloud messaging, and analytics
- **Supabase**: Storage and real-time database
- **SendGrid**: Email service for notifications and communication

### Tracking & Sensors

- **Google Maps API**: Location tracking and visualization
- **Sensors Plus**: Fall detection using device sensors
- **Geolocator**: GPS location tracking

## 🚀 Installation and Setup

### Prerequisites

- Go 1.18 or newer
- Docker and Docker Compose
- PostgreSQL (optional for local development without Docker)
- Git

### Installation Steps

1. **Clone the repository**

   ```
   git clone https://github.com/elginbrian/ELDERWISE-BE.git
   ```

2. **Navigate to the project directory**

   ```
   cd ELDERWISE-BE
   ```

3. **Set up environment variables**

   Create a `.env` file based on the examples in docker-compose files:

   ```
   # Database settings
   POSTGRES_HOST=localhost
   POSTGRES_PORT=5432
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=postgres
   POSTGRES_DB=elderwise_dev
   POSTGRES_TIMEZONE=Asia/Jakarta

   # Storage settings
   SUPABASE_URL=your_supabase_url
   SUPABASE_KEY=your_supabase_key
   SUPABASE_BUCKETNAME=elderwise-images

   # Email settings
   SENDGRID_API_KEY=your_sendgrid_key
   EMAIL_FROM=your_email@example.com
   EMAIL_FROM_NAME=Elderwise Alert System
   ```

4. **Install dependencies**

   ```
   go mod tidy
   go mod vendor
   ```

5. **Run with Docker (recommended)**

   For development:

   ```
   docker-compose -f docker-compose.dev.yml up -d
   ```

   For production:

   ```
   docker-compose -f docker-compose.prod.yml up -d
   ```

6. **Run without Docker**

   Make sure PostgreSQL is running locally, then:

   ```
   go run cmd/elderwise/main.go
   ```

7. **Access the API**

   The API will be available at:

   - Development: http://localhost:4000/api/v1
   - Production: http://localhost:4001/api/v1

## 👥 Team Masukin Andre ke Raion

- **Andreas Bagasgoro** - _PM and UI/UX_
- **Muhammad Rizqi Aditya Firmansyah** - _Frontend and UI/UX_
- **Elgin Brian Wahyu Bramadhika** - _Backend and Frontend_

## 🏆 FIND-IT UGM 2025 Hackathon

This project was developed for the FIND-IT UGM 2025 Hackathon competition, focusing on technological solutions to improve the quality of life for elderly people through accessible and user-friendly innovations.

---
