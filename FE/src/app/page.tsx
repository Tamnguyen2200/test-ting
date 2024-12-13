"use client"

import React, {useEffect, useRef, useState} from 'react';
// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'

interface Kols {
	KolID: number;
	UserProfileID: number;
	Language: string;
	Education: string;
	ExpectedSalary: number;
	ExpectedSalaryEnable: boolean;
	ChannelSettingTypeID: number;
	IDFrontURL: string;
	IDBackURL: string;
	PortraitURL: string;
	RewardID: number;
	PaymentMethodID: number;
	TestimonialsID: number;
	VerificationStatus: boolean;
	Enabled: boolean;
	ActiveDate: Date;
	Active: boolean;
	CreatedBy: string;
	CreatedDate: Date;
	ModifiedBy: string;
	ModifiedDate: Date;
	IsRemove: boolean;
	IsOnBoarding: boolean;
	Code: string;
	PortraitRightURL: string;
	PortraitLeftURL: string;
	LivenessStatus: boolean;
}

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [Kols , setKols] = useState<Kols[]>([]);  

    // Loading state
    const [loading, setLoading] = useState<boolean>(false); 

    // Pagination state
    const [currentPage, setCurrentPage] = useState<number>(1);
    const [itemsPerPage, setItemsPerPage] = useState<number>(5);
    const [totalCount, settotalCount] = useState<number>(0);


    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
    useEffect(() => {
        async function fetchKols() {
            setLoading(true);
            try {
                const response = await fetch(`http://localhost:8081/kols?pageIndex=${currentPage}&pageSize=${itemsPerPage}`); 
                const data = await response.json();
                if(data.result == 'Success'){
                    setKols(data.kol);
                    settotalCount(data.totalCount)
                } else {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                
            } catch (error) {
                console.error('Error fetching Kol data:', error);
            } finally {
                setLoading(false); // Set loading to false after the request completes
            }
        }
        fetchKols();
    }, [itemsPerPage, currentPage]);

  
    // Calculate total pages
    const totalPages = Math.ceil(totalCount / itemsPerPage);
  
    // Handle page change
    const handlePageChange = (pageNumber: number) => {
      setCurrentPage(pageNumber);
    };

	const handleItemsPerPageChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setItemsPerPage(Number(event.target.value));
        setCurrentPage(1); // Reset to first page when items per page is changed
    };

    return (
        <>
			<div className='Container'>
				<table className="table">
					<thead>
						<tr>
							<th>Kol ID</th>
							<th>User Profile ID</th>
							<th>Language</th>
							<th>Education</th>
							<th>Expected Salary</th>
							<th>Expected Salary Enable</th>
							<th>Channel Setting Type ID</th>
							<th>ID Front URL</th>
							<th>ID Back URL</th>
							<th>Portrait URL</th>
							<th>Reward ID</th>
							<th>Payment Method ID</th>
							<th>Testimonials ID</th>
							<th>Verification Status</th>
							<th>Enabled</th>
							<th>Active Date</th>
							<th>Active</th>
							<th>Created By</th>
							<th>Created Date</th>
							<th>Modified By</th>
							<th>Modified Date</th>
							<th>Is Remove</th>
							<th>Is On Boarding</th>
							<th>Code</th>
							<th>Portrait Right URL</th>
							<th>Portrait Left URL</th>
							<th>Liveness Status</th>
						</tr>
					</thead>
					{loading ? ( 
						<div className="loading-spinner">
							<div className="spinner"></div> {/* Simple spinner */}
						</div>
					) : (
					<tbody>
					{Kols.map((kolData) => (
						<tr>
							<td>{kolData.KolID}</td>
							<td>{kolData.UserProfileID}</td>
							<td>{kolData.Language}</td>
							<td>{kolData.Education}</td>
							<td>{kolData.ExpectedSalary}</td>
							<td>{kolData.ExpectedSalaryEnable ? 'Enabled' : 'Disabled'}</td>
							<td>{kolData.ChannelSettingTypeID}</td>
							<td><a href={kolData.IDFrontURL} target="_blank" rel="noopener noreferrer">View</a></td>
							<td><a href={kolData.IDBackURL} target="_blank" rel="noopener noreferrer">View</a></td>
							<td><a href={kolData.PortraitURL} target="_blank" rel="noopener noreferrer">View</a></td>
							<td>{kolData.RewardID}</td>
							<td>{kolData.PaymentMethodID}</td>
							<td>{kolData.TestimonialsID}</td>
							<td>{kolData.VerificationStatus ? 'Verified' : 'Pending'}</td>
							<td>{kolData.Enabled ? 'Enabled' : 'Disabled'}</td>
							<td>{new Date(kolData.ActiveDate).toLocaleDateString()}</td>
							<td>{kolData.Active ? 'Active' : 'Inactive'}</td>
							<td>{kolData.CreatedBy}</td>
							<td>{new Date(kolData.CreatedDate).toLocaleDateString()}</td>
							<td>{kolData.ModifiedBy}</td>
							<td>{new Date(kolData.ModifiedDate).toLocaleDateString()}</td>
							<td>{kolData.IsRemove ? 'Yes' : 'No'}</td>
							<td>{kolData.IsOnBoarding ? 'Yes' : 'No'}</td>
							<td>{kolData.Code}</td>
							<td><a href={kolData.PortraitRightURL} target="_blank" rel="noopener noreferrer">View</a></td>
							<td><a href={kolData.PortraitLeftURL} target="_blank" rel="noopener noreferrer">View</a></td>
							<td>{kolData.LivenessStatus ? 'Passed' : 'Failed'}</td>
						</tr>
					))}
					</tbody>
					)}
				</table>
				
				<div className="pagination">
					<button 
						onClick={() => handlePageChange(currentPage - 1)} 
						disabled={currentPage === 1}
						>
						Previous
					</button>

					{/* Display page numbers */}
					{[...Array(totalPages)].map((_, index) => (
						<button
							key={index + 1}
							onClick={() => handlePageChange(index + 1)}
							className={currentPage === index + 1 ? 'active' : ''}
						>
							{index + 1}
						</button>
					))}

					<button 
						onClick={() => handlePageChange(currentPage + 1)} 
						disabled={currentPage === totalPages}
						>
						Next
					</button>

					<div className="items-per-page">
						<label htmlFor="itemsPerPage">Items per page:</label>
						<select
						id="itemsPerPage"
						value={itemsPerPage}
						onChange={handleItemsPerPageChange}
						>
							<option value={5}>5</option>
							<option value={10}>10</option>
							<option value={15}>15</option>
							<option value={20}>20</option>
						</select>
					</div>
				</div>
			</div>        
		</>
    )
};

export default Page;