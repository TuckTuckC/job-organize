# 🛠️ Project Notes: [Job-Organize]

> A full solution to organization for small businesses.

---

🧰 Technologies

- 🖥️  Frontend:
  - HTML
  - CSS
  - JavaScript
  - Tailwind.css

- 🧪 Backend:
  - Rust 🦀

## ✅ Current Features (In Development)

- [ ] **Employee Clock In/Out**
  - Track time for active shifts
  - Option for companies to track location
  - Timesheet Analasys

- [ ] **Calander / Scheduling System**
  - Will display items from Task List on that day

- [ ] **PTO Requests**
  - Employees can request time off
  - Manual or admin-approved workflow TBD
    
## 🚧 TODOs

- [ ] Decide on **project name**
- [ ] Define minimum feature set for v1
- [ ] Create mock UI wireframes
- [ ] Set up database schema
- [ ] Draft roles (admin, employee, manager, etc.)

## 🗃️ Database Structure
- Users
  - Personal Info
  - Role
- Companies
  - Users (fk)
  - Crews
  - Job Locations
    - Current Priority Task List
    - Inventory List
    - Notes
    - Additional Files
  - Clients
  - Inventory List

## 💡 Future Ideas

- **Automation**
    - Automatic Billing
    - Send reminders via push notifications
    - Schedule follow-ups

- **Reporting & Insights**
  - Payroll summaries
  - Job completion stats
  - PTO analysis

- **Mobile Accessibility**
  - Responsive interface
  - Consider push notifications or location-aware clocking

## 🧪 Experimental

> Use this section for notes, testing, or brainstorming

- Thoughts on **webhooks** for triggering automation?
- Manual override for billing in case of client changes?
- File uploads for job documentation (images, PDFs)?

